package routers

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
	"url-shortner/pkg/logger"
)

const (
	RUNNING = iota
	TERMINATING
)

// Atomic struct encapsulating server state
type ServerState struct {
	running int32
}

// Atomically store that the server is shutting down
func (state *ServerState) Shutdown() {
	atomic.StoreInt32(&state.running, TERMINATING)
}

// Atomically check that is running
func (state *ServerState) IsRunning() bool {
	return atomic.LoadInt32(&state.running) == RUNNING
}

// type Graceful will run an http.Server, with graceful shutdown
type Graceful struct {
	Server          *http.Server
	ShutdownTimeout time.Duration
	State           *ServerState
}

// ListenAndServe the enclosed http.Server, but shutdown gracefully
func (g *Graceful) ListenAndServe() {
	logger.Info("Listening and serving on port ", g.Server.Addr)

	go g.Server.ListenAndServe()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal)

	// Register our signal handlers in order
	// ^C in the terminal will send an os.Interrupt
	signal.Notify(quit, os.Interrupt)
	// Kubernetes will send a SIGTERM, so notify on that as well
	signal.Notify(quit, syscall.SIGTERM)

	foundSignal := <-quit
	g.State.Shutdown()

	if foundSignal == syscall.SIGTERM {
		// If we terminate immediately from a SIGTERM, we still may
		// have incoming connections routed to us by k8s. Instead,
		// disable KeepAlives, and sleep to wait for this to propagate.

		g.Server.SetKeepAlivesEnabled(false)

		logger.Info("SIGTERM received, starting to shut down")
		time.Sleep(15 * time.Second)
	}

	logger.Info("Gracefully shutting down server with timeout: ", g.ShutdownTimeout)

	ctx, cancel := context.WithTimeout(context.Background(), g.ShutdownTimeout)
	defer cancel()
	if err := g.Server.Shutdown(ctx); err != nil {
		logger.Fatal("Server Shutdown: ", err)
	}

	logger.Info("Server exiting")
	os.Exit(0)
}

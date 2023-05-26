package routers

import (
	"fmt"
	"net/http"
	"url-shortner/pkg/config"
	"url-shortner/pkg/logger"
)

func runServer(app config.AppConfig) {

	server := &http.Server{
		Addr:    ":" + app.Config.Server.Port,
		Handler: app.Router,
	}

	displayService()

	logger.Infof("Serving and listening on Port %s ", server.Addr)
	server.ListenAndServe()
}

// You can generate ASCI art here
// https://patorjk.com/software/taag/#p=display&f=Doom&t=sample
func displayService() {
	fmt.Println(`
	
	_____ __ __   ___   ____  ______      __ __  ____   _           ____    ___  ____     ___  ____    ____  ______   ___   ____  
	/ ___/|  |  | /   \ |    \|      |    |  |  ||    \ | |         /    |  /  _]|    \   /  _]|    \  /    ||      | /   \ |    \ 
   (   \_ |  |  ||     ||  D  )      |    |  |  ||  D  )| |        |   __| /  [_ |  _  | /  [_ |  D  )|  o  ||      ||     ||  D  )
	\__  ||  _  ||  O  ||    /|_|  |_|    |  |  ||    / | |___     |  |  ||    _]|  |  ||    _]|    / |     ||_|  |_||  O  ||    / 
	/  \ ||  |  ||     ||    \  |  |      |  :  ||    \ |     |    |  |_ ||   [_ |  |  ||   [_ |    \ |  _  |  |  |  |     ||    \ 
	\    ||  |  ||     ||  .  \ |  |      |     ||  .  \|     |    |     ||     ||  |  ||     ||  .  \|  |  |  |  |  |     ||  .  \
	 \___||__|__| \___/ |__|\_| |__|       \__,_||__|\_||_____|    |___,_||_____||__|__||_____||__|\_||__|__|  |__|   \___/ |__|\_|
																																   
   
		`)
}

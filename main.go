package main

import (
	"url-shortner/pkg/config"
	"url-shortner/pkg/database"
	"url-shortner/routers"

	"url-shortner/pkg/logger"
)

func main() {
	//reading config from env
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("unable to load config ", err)
	}
	db, err := database.GetDB(cfg)
	if err != nil {
		logger.Fatal("unable to load config ", err)
	}
	app := config.AppConfig{
		DB:     db,
		Config: cfg,
	}

	routers.SetupAndRunServer(&app)
}

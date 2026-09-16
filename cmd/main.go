package main

import (
	"log/slog"
	"os"
	"pizza-tracker-go/internal/models"
	"github.com/gin-gonic/gin"


)

func main() {
	cfg := loadConfig()

	logger :=slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	//  seting the data base
	dbModel, err := models.InitDB(cfg.DBSource)
	if err != nil {
		slog.Error("failed to initialize database", "error", err)
		os.Exit(1)
	}

	slog.Info("Database initialized sucessfully")

	RegisterCustomValidators()

	h:= NewHandler(dbModel)

	router := gin.Default()
	if err != nil {
		slog.Error("Failed to load templates", "error", err)
		os.Exit(1)
	}

	setupRoutes(router, h)

	slog.Info("Starting server", "url", "http://localhost:"+cfg.Port, "and port no", cfg.Port)

	router.Run(":"+cfg.Port)

}
package main

import (
	"context"
	"fmt"
	"net/http"
	"shop-bot/internal/common/logger"
	"shop-bot/internal/common/middleware"
	"shop-bot/internal/config"
	"shop-bot/internal/database/migrations"
	"shop-bot/internal/delivery/telegram"
	"shop-bot/internal/delivery/telegram/utils"
)

func main() {
	log := logger.NewLogger()

	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName, cfg.Database.SSLMode,
	)

	err = migrations.RunMigrations(dsn)
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	router := telegram.NewRouter()

	router.RegisterCommand("start", telegram.StartCommand)
	router.RegisterCommand("help", telegram.HelpCommand)
	router.RegisterCommand("order", telegram.HandleOrder)
	router.RegisterCommand("contacts", telegram.HandleContacts)

	botHandler, err := telegram.NewBotHandler(&utils.TelegramConfig{
		BotToken:            cfg.Bot.Token,
		EnableMultilanguage: cfg.Bot.EnableMultilanguage,
		DefaultLanguage:     cfg.Bot.DefaultLanguage,
		AllowedLanguages:    cfg.Bot.AllowedLanguages,
	}, router)
	if err != nil {
		log.Fatalf("failed to initialize bot handler: %v", err)
	}

	go func() {
		if err := botHandler.Start(); err != nil {
			log.Fatalf("Telegram bot error: %v", err)
		}
	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Shop Bot API is running"))
	})

	loggedRouter := middleware.LoggingMiddleware(log.Logger)(mux)
	middleware.RecoveryMiddleware(log.Logger)(loggedRouter)
}

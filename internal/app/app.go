package app

import (
	"log"
	"log/slog"
	"os"

	"github.com/BeRebornBng/OsauAmsApi/internal/config"
	"github.com/BeRebornBng/OsauAmsApi/internal/handler"
	"github.com/BeRebornBng/OsauAmsApi/internal/repository"
	"github.com/BeRebornBng/OsauAmsApi/internal/server"
	"github.com/BeRebornBng/OsauAmsApi/internal/service"
	"github.com/BeRebornBng/OsauAmsApi/pkg/auth"
	"github.com/BeRebornBng/OsauAmsApi/pkg/database/postgres"
	"github.com/BeRebornBng/OsauAmsApi/pkg/myhash"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

const configPath = "configs"

func Run() {
	// TO DO CONFIG
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Fatal("failed to init config", err)
	}

	// TO DO INIT LOGGER
	log := setupLogger(cfg.Env)
	log.Info(
		"starting osau ams api",
		slog.String("env", cfg.Env),
	)

	// TO DO INIT DATABASE
	db, err := postgres.New(cfg.Postgres.Url)
	if err != nil {
		log.Debug("unable to connect to database: ", err)
		os.Exit(1)
	}
	defer func() {
		db.Close()
		log.Debug("database connection closed")
	}()
	hasher := myhash.NewHasher("salt")
	tokenManager := auth.NewManager(cfg.Jwt.SecretKey)
	if err != nil {
		log.Debug(err.Error())
	}

	// TO DO INIT REPOSITORIES
	repos := repository.NewRepositories(db)

	// TO DO INIT SERVICES
	services := service.NewServices(
		service.Support{
			Repos:          repos,
			Hasher:         hasher,
			TokenManager:   tokenManager,
			AccessTokenTTL: cfg.Jwt.AccessTokenTTL,
		},
	)

	// TO DO INIT ROUTER
	h := handler.NewHandler(tokenManager, services, log)

	// TO DO RUN SERVER
	s := server.NewServer(cfg, h.InitRoutes())
	if err := s.Run(); err != nil {
		log.Debug(err.Error())
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}

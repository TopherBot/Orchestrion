package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/yourorg/orchestrion/internal/api"
	"github.com/yourorg/orchestrion/internal/engine"
)

func initConfig() {
	viper.SetDefault("server.port", "8080")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // cwd
	_ = viper.ReadInConfig() // ignore error; env overrides are fine
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	initConfig()

	port := viper.GetString("server.port")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// API routes
	api.RegisterRoutes(router)

	// Health endpoint
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	// Initialize engine (worker pool, SBOM, etc.) – simplified
	if err := engine.Start(context.Background()); err != nil {
		log.Fatal().Err(err).Msg("failed to start engine")
	}

	addr := ":" + port
	log.Info().Msgf("Orchestrion listening on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal().Err(err).Msg("server stopped")
	}
}

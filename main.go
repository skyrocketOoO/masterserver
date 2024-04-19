package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	errors "github.com/rotisserie/eris"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/skyrocketOoO/masterserver/api"
	"github.com/skyrocketOoO/masterserver/config"
	"github.com/skyrocketOoO/masterserver/internal/delivery/rest"
	"github.com/skyrocketOoO/masterserver/internal/delivery/rest/middleware"
	"github.com/skyrocketOoO/masterserver/internal/infra/postgres"
	"github.com/skyrocketOoO/masterserver/internal/usecase"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	// human-friendly logging without efficiency
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Logger initialized")

	if err := config.ReadConfig(); err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}

	db, disconnectDb, err := postgres.InitDB()
	if err != nil {
		log.Fatal().Msg(errors.ToString(err, true))
	}
	defer disconnectDb()
	dbRepo := postgres.NewOrmRepository(db)

	usecase := usecase.NewUsecase(dbRepo)
	delivery := rest.NewRestDelivery(usecase)

	router := gin.Default()
	router.Use(middleware.CORS())
	api.Binding(router, delivery)

	router.Run(":8081")
}

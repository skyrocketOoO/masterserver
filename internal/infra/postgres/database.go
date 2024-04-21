package postgres

import (
	"fmt"

	errors "github.com/rotisserie/eris"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() (*gorm.DB, func(), error) {
	log.Info().Msg("Connecting to Postgres")
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.db"),
		viper.GetString("postgres.sslmode"),
		viper.GetString("postgres.timezone"),
	)

	db, err := gorm.Open(
		postgres.Open(connStr), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		},
	)
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}

	var disconnect = func() {
		db, _ := db.DB()
		db.Close()
	}
	return db, disconnect, nil
}

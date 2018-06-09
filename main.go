package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/spf13/viper"

	"github.com/picobank/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	log15 "gopkg.in/inconshreveable/log15.v2"
)

//go:generate sqlboiler postgres

var log = log15.New()

func main() {
	var err error

	// Load configuration
	err = initViper()
	if err != nil {
		fmt.Println("unable to load config file")
		os.Exit(-2)
	}

	db, err := sql.Open("pgx", createConnectionString())
	if err != nil {
		log.Error("Failed to open database", "err", err)
		return
	}
	log.Info("Connected to DB.")

	instrument, err := models.Instruments(db, qm.Limit(1)).One()
	if err != nil {
		log.Error("Failed to select Instrument", "err", err)
		return
	}
	log.Info("Fetched instrument", "id", instrument.InstrumentID)
}

func initViper() error {
	// TODO: This is rudimentary. Has to be improved...
	log.Debug("Reading configuration...")
	viper.SetConfigName("sqlboiler")
	wd, err := os.Getwd()
	if err != nil {
		log.Error("Error getting current diretory", "err", err)
		return err
	}
	viper.AddConfigPath(wd)
	err = viper.ReadInConfig()
	if err != nil {
		log.Error("Unable to read configuration file", "err", err)
		return err
	}
	viper.AutomaticEnv()

	return nil
}

func createConnectionString() string {
	return fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s",
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.host"),
		viper.GetString("postgres.port"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.password"),
		viper.GetString("postgres.sslmode"),
	)
}

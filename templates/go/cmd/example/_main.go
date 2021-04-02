package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"{{ module }}/cmd/fulfillment/utils"
	"{{ module }}/http"
	"{{ module }}/sql/postgres"
)

func main() {

	var cfg utils.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var environment = cfg.Environment
	var logLevel = cfg.LogLevel
	initLogger(logLevel)

	db, err := postgres.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// inject database into services
	exampleService := postgres.NewExampleService(db)

	httpserver := http.NewServer(environment)
	httpserver.ExampleService = exampleService

	if err := httpserver.Open(); err != nil {
		log.Fatalf("error: %s", err)
	}

	defer httpserver.Close()

}

func initLogger(level string) {
	parsedLevel, err := log.ParseLevel(level)
	if err != nil {
		log.Error(err)
		parsedLevel = log.InfoLevel
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(parsedLevel)
}

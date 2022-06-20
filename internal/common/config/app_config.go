package config

import (
	"fmt"
	"strings"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

func App() appConfig {
	ensureConfigLoaded()
	return app
}

var app appConfig

type appConfig struct {
	ENV     string
	Port    int
	Domains []string
	Secret  string
}

func loadAppConfig() {
	app = appConfig{
		ENV:     getENV("APP_ENV"),
		Port:    getIntENV("APP_PORT"),
		Domains: strings.Split(getENV("APP_DOMAINS"), ";"),
	}

	if !(app.ENV == "dev" || app.ENV == "staging" || app.ENV == "prod") {
		panic(fmt.Sprintf("Expected env with key 'APP_ENV' to be '%s' or '%s', found '%v'", app.ENV, DEV, PROD))
	}

	if app.Domains == nil || len(app.Domains) == 0 {
		if app.ENV == "prod" {
			panic("APP_DOMAINS env variable must be set")
		}
		app.Domains = []string{"localhost"}
	}
}

package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

const appPrefix = "APP"

var App AppSpec

type AppSpec struct {
	Port int `envconfig:"PORT" default:"1323"`
}

func init() {
	log.Println("initialize app config...")
	err := envconfig.Process(appPrefix, &App)
	if err != nil {
		log.Fatal(err.Error())
	}
}

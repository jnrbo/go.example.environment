package config

import (
	"fmt"
	"go.example.environment/app/config/development"
	"go.example.environment/app/config/production"
	"os"
)

var databaseSources = map[string]func() string {
	"production": production.GetSource,
	"development": development.GetSource,
}

func Environment() string {
	env, found := os.LookupEnv("ENVIRONMENT")
	if found {
		fmt.Print("Enviroment: " + env)
		return env
	} else {
		return "development"
	}
}

func GetSource() string {
	 return databaseSources[Environment()]()
}

func IsProduction() bool {
	return Environment() == "production"
}
package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Port        int
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}

func GetConfig(configType string) (*Config, error) {
	switch configType {
	case "env":
		return getEnvConfig()
	default:
		return nil, fmt.Errorf("type \"%s\" is unsupported", configType)
	}
}

// getEnvConfig parse env-vars and return *config.Config type.
func getEnvConfig() (*Config, error) {
	config := &Config{}

	port, err := getEnvVar("PORT")
	if err != nil {
		return nil, err
	}
	portInt, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		return nil, err
	}
	config.Port = int(portInt)

	dbUrl, err := getEnvVar("DB_URL")
	if err != nil {
		return nil, err
	}
	config.DbUrl = dbUrl

	jaegerUrl, err := getEnvVar("JAEGER_URL")
	if err != nil {
		return nil, err
	}
	config.JaegerUrl = jaegerUrl

	sentryUrl, err := getEnvVar("SENTRY_URL")
	if err != nil {
		return nil, err
	}
	config.SentryUrl = sentryUrl

	kafkaBroker, err := getEnvVar("KAFKA_BROKER")
	if err != nil {
		return nil, err
	}
	config.KafkaBroker = kafkaBroker

	someAppId, err := getEnvVar("SOME_APP_ID")
	if err != nil {
		return nil, err
	}
	config.SomeAppId = someAppId

	someAppKey, err := getEnvVar("SOME_APP_KEY")
	if err != nil {
		return nil, err
	}
	config.SomeAppKey = someAppKey

	return config, nil
}

func getEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		return "", fmt.Errorf("you should define env-var \"%s\"", envVar)
	}
	return value, nil
}

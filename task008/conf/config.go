package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Port        int    `yaml:"port"`
	DbUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppId   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

func GetConfig(configType string) (*Config, error) {
	switch {
	case configType == "env":
		return getEnvConfig()
	case strings.HasSuffix(configType, ".yml"):
		return getYamlConfig(configType)
	default:
		return nil, fmt.Errorf("type \"%s\" is unsupported", configType)
	}
}

func getYamlConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatalf("couldn't close the file: %s", filepath)
		}
	}()

	config := &Config{}
	err = yaml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
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

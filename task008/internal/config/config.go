package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Config struct {
	Port        int    `env-var:"PORT"`
	DbUrl       string `env-var:"DB_URL"`
	JaegerUrl   string `env-var:"JAEGER_URL"`
	SentryUrl   string `env-var:"SENTRY_URL"`
	KafkaBroker string `env-var:"KAFKA_BROKER"`
	SomeAppId   string `env-var:"SOME_APP_ID"`
	SomeAppKey  string `env-var:"SOME_APP_KEY"`
}

func getEnvVar(envVar string) (string, error) {
	value := os.Getenv(envVar)
	if value == "" {
		return "", fmt.Errorf("you should define env-var \"%s\"", envVar)
	}
	return value, nil
}

// GetConfigByReflect parse env-vars and write values into config.
// config should be a pointer on a struct.
// Only string and int fields for "config" have been supported.
func GetConfigByReflect(config interface{}) error {

	objectType := reflect.TypeOf(config)
	objectValue := reflect.ValueOf(config)

	// Проходим по всем полям структуры
	for i := 0; i < objectType.Elem().NumField(); i++ {
		// Если у поля нет тега "env-var", то пропускаем поле
		envVarName, ok := objectType.Elem().Field(i).Tag.Lookup("env-var")
		if !ok {
			continue
		}

		// Получаем значение переменной окружения
		envVarValue, err := getEnvVar(envVarName)
		if err != nil {
			return err
		}

		// Записываем значение переменной окружения в поле структуры
		if objectType.Elem().Field(i).Type.Kind() == reflect.Int {
			intValue, err := strconv.ParseInt(envVarValue, 10, 64)
			if err != nil {
				return err
			}
			objectValue.Elem().Field(i).SetInt(intValue)
			continue
		}
		objectValue.Elem().Field(i).SetString(envVarValue)
	}

	return nil
}

// GetConfig parse env-vars and return *config.Config type.
func GetConfig() (*Config, error) {
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

package main

import (
	"env-var/internal/config"
	"fmt"
	"log"
)

func main() {
	conf := &config.Config{}

	// Используем рефлексию для получения переменных окружения.
	// Использовать рефлексию, когда можно обойтись без неё - не best practice.
	// Но знать как использовать определенно стоит
	err := config.GetConfigByReflect(conf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)

	// Без использования рефлексии
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
}

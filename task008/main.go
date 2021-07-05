package main

import (
	"env-var/config"
	"flag"
	"fmt"
	"log"
)

var (
	configType = flag.String("type", "env", "Source for getting config. Options: env")
)

func main() {
	flag.Parse()

	err, conf := config.GetConfig(*configType)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)
}

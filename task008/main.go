package main

import (
	"env-var/conf"
	"flag"
	"fmt"
	"log"
)

var (
	config = flag.String("type", "env", "Source for getting config. \nUse 'env' for getting values from env-vars.\nUse 'filepath.yml' for getting values from yaml-file")
)

func main() {
	flag.Parse()

	conf, err := conf.GetConfig(*config)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(conf)
}

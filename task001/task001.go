package main

import "fmt"

func main() {
	if _, err := fmt.Println("Hello, World!"); err != nil {
		panic(err)
	}
}

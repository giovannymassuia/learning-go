package main

import "github.com/giovannymassuia/learning-go/07-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBHost)
}

package main

import (
	"fmt"

	utils "github.com/mikevel2955/hermes-utils"
)

type config struct {
	TestStr         string `env:"TEST_STR" def:"Hello world!!!"`
	TestStrRequired string `env:"TEST_STR_REQUIRED" required:"true"`
}

func main() {
	config := config{}
	if err := utils.ReadConfig(&config); err != nil {
		panic(err)
	}
	fmt.Println("config:", config)
}

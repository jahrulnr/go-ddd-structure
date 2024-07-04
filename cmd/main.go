package main

import (
	"ddd-impl/config"
	"ddd-impl/internal/infrastructure/delivery"
	"fmt"
	"log"
	"os"
)

func main() {
	test := os.Getenv("APP_ON_PRIME")
	fmt.Println(test)
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	delivery.Run(cfg)
}

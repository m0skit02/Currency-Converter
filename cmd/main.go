package main

import (
	converter "CurrencyConverter"
	"log"
)

func main() {
	srv := new(converter.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error %s", err.Error())
	}
}

package main

import (
	"github.com/joho/godotenv"
	"github.com/oasis-prime/oas-platform-hotels-master-api/protocol"
)

func main() {
	godotenv.Load()

	err := protocol.ServeHTTP()
	if err != nil {
		panic(err)
	}
}

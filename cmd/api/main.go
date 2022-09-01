package main

import (
	"github.com/oasis-prime/oas-platform-hotels-master-api/protocol"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		panic(err)
	}
}

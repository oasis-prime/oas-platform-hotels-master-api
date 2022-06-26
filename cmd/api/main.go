package main

import (
	"fmt"

	"github.com/oasis-prime/oas-platform-hotels-master-api/protocol"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		fmt.Println(err)
	}
}

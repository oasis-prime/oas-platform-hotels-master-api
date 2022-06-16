package main

import (
	"fmt"
	"ss-platform-api/protocol"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"os"
	"fmt"
	"net"
	"strings"
	"log"
)

func main() {

	arguments := os.Args

	for _, zoneAddress := range arguments {

		if !strings.Contains(zoneAddress, "=") {
			continue
		}

		keyValue := strings.Split(zoneAddress, "=")

		if len(keyValue) != 2 {
			log.Fatalf("Input does not match expectations. Expected: Zone=Address, e.g. London=localhost:8080")
			os.Exit(1)
		}

		zone := keyValue[0]
		address := keyValue[1]

		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatalf("Unable to connect to \"%s\"", address)
			os.Exit(1)
		}
		defer conn.Close()


	}

}
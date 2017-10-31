package main

import (
	"os"
	"net"
	"strings"
	"log"
	"fmt"
	"time"
	"io"
)

func main() {

	arguments := os.Args

	zoneConnection := make(map[string]net.Conn)

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

		conn.SetReadDeadline(time.Time{})
		zoneConnection[zone] = conn
	}

	for {
		for zone, conn := range zoneConnection {
			readFromClockServer(zone, conn)
		}
		time.Sleep(1 * time.Second)
		fmt.Printf("\r")
	}
}

func readFromClockServer(zone string, conn net.Conn) {
	buffer := make([]byte, 9)
	if _, err := io.ReadAtLeast(conn, buffer, 9); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s:\t%s \t", zone, strings.Trim(string(buffer), "\n"))
}

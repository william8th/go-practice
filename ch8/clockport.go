package main

import (
	"net"
	"log"
	"io"
	"time"
	"flag"
	"fmt"
	"os"
)

func main() {

	var port int
	var timezone string
	flag.IntVar(&port, "port", -1, "The port number to run the clock from")
	flag.StringVar(&timezone, "timezone", "Europe/London", "The time zone of the clock")

	flag.Parse()

	if port < 0 {
		log.Printf("Port input is mandatory. Cannot run clock server at port %d, please supply a valid port number", port)
		os.Exit(1)
	}

	address := fmt.Sprintf("localhost:%d", port)

	location, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Printf("Running clock server at %s with time zone '%s'", address, timezone)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn, location)
	}
}

func handleConn(c net.Conn, location *time.Location) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().In(location).Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

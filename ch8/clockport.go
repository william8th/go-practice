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
	flag.IntVar(&port, "port", -1, "The port number to run the clock from")

	flag.Parse()

	if port < 0 {
		log.Printf("Port input is mandatory. Cannot run clock server at port %d, please supply a valid port number", port)
		os.Exit(1)
	}

	address := fmt.Sprintf("localhost:%d", port)

	log.Printf("Running server at %s", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

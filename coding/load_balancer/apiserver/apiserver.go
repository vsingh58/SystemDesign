package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {

	port := flag.Int("port", 8080, "tcp port")
	flag.Parse()

	address := fmt.Sprintf(":%d", *port)
	fmt.Println("Listening on address:", address)
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Failed to bind to port")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Received processing request..")
		go doProcessHeader(conn)
	}
}

func doProcessHeader(conn net.Conn) {

	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		return
	}
	payload := string(buf[:n])
	fmt.Sprintf("Received: %s\n", payload)
}

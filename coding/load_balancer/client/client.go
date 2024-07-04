package main

import (
	"fmt"
	"net"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	go doProcess(conn)
}

func doProcess(conn net.Conn) {
	defer conn.Close()

	// Send data to the server
	_, write_err := conn.Write([]byte("Hello, server!"))
	if write_err != nil {
		fmt.Println(write_err)
		return
	}

	// Read and process data from the server
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the incoming data
	fmt.Printf("Received: %s", buf)

}

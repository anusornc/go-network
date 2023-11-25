package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close() // close connection before exit

	// buffer for reading
	buffer := make([]byte, 1024)
	for {
		// Read data from the client
		n, err := conn.Read(buffer)	// Read() blocks until it reads some data from the network and n is the number of bytes read
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}
		// Print the number of bytes read
		fmt.Printf("Received %d bytes\n", n)

		// Print received data
		// [0 : n] , [:n] , [0:n] are all equivalent
		fmt.Printf("Received message: %s", buffer[:n]) // :n is a slice operator that returns a slice of the first n bytes of the buffer

		// Print the message as bytes
		fmt.Printf("Received message as bytes: %v\n", buffer[:n])

		// 104 is the ASCII code for 'h'
		// 105 is the ASCII code for 'i'
		// 10 is the ASCII code for '\n'

		// Send a response back to the client
		response := "Message received successfully\n"
		conn.Write([]byte(response))
	}
}

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", ":5000")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}

	// Close the listener when the application closes
	defer listener.Close()

	fmt.Println("Server is listening on port 5000")

	// Listen for an incoming connection
	for {
		// Accept() blocks until a connection is made
		conn, err := listener.Accept() // Accept is three-way handshake
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue // continue to next iteration of for loop
		}

		fmt.Println("New connection established")

		// Handle connections in a new goroutine
		go handleConnection(conn)
	}
}
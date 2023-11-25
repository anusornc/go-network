package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()	// Close connection before exiting the function

	// Create a buffer to store received data
	// The buffer size is 1024 bytes
	// If the received data is larger than 1024 bytes, the buffer will be filled with the first 1024 bytes
	// The remaining bytes will be discarded
	// If the received data is smaller than 1024 bytes, the buffer will be filled with the received data
	// The remaining bytes will be filled with zero values
	// The buffer is a slice of bytes, so it can be converted to a string using the string() function
	// what if data is larger than 1024 bytes?
	// we can use a loop to read data in chunks of 1024 bytes
	// we can use a dynamic buffer that grows as needed
	// we can use a buffer pool to reduce memory allocations
	// why 1024 bytes?
	// it's a reasonable size for a buffer
	// it's a power of 2, so it's more efficient to allocate and use it in memory (1024 = 2^10)
	// it's a multiple of the TCP maximum segment size (MSS) which is 1460 bytes by default (1460 * 7 = 10220) 7 is the maximum number of segments that can be sent in a single TCP window
	// it's a multiple of the Ethernet MTU which is 1500 bytes by default )
	// it's a multiple of the IPv4 MTU which is 576 bytes by default
	// can we use a buffer size smaller than 1024 bytes?
	// yes, but it will be less efficient because we will need to read more data from the network
	// can we use a buffer size larger than 1024 bytes?
	// yes, but it will be less efficient because we will need to allocate more memory
	// if we send 'hello' which is 5 bytes, the buffer will be filled with 'hello' and the remaining bytes will be filled with zero values
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
		fmt.Printf("Received message: %s", buffer[:n]) // :n is a slice operator that returns a slice of the first n bytes of the buffer

		// Send a response back to the client
		response := "Message received successfully\n"
		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("New connection established")

		go handleConnection(conn)
	}
}

package main

import (
	"net"
)

func handleConnection(connection net.Conn) {
	defer connection.Close()

	buffer := make([]byte, 1024)
	for {
		idx, err := connection.Read(buffer)
		if err != nil {
			println("Error reading:", err.Error())
			return
		}

		println("Received message:", string(buffer[:idx]))
		_, err = connection.Write(buffer[:idx])
		if err != nil {
			println("Error writing:", err.Error())
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		println("Error listening:", err.Error())
		return
	}
	defer listen.Close()
	println("Listening on port 8080")

	for {
		connection, err := listen.Accept()
		println("Accepted connection")
		if err != nil {
			println("Error accepting:", err.Error())
			continue
		}

		go handleConnection(connection)
	}
}

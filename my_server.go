package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// closing the connection
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Failed to close connection")
			return
		}
		fmt.Println("Connection closed successfully")
	}(conn)

	reader, writer := bufio.NewReader(conn), bufio.NewWriter(conn)

	for {
		//Read data from the connection
		data, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println("Failed to read data: ", err.Error())
			return
		}
		fmt.Println("Data read successfully", string(data[:]))

		// sending the client an ack
		_, err = writer.Write(data)
		if err != nil {
			fmt.Println("Failed to write data: ", err.Error())
			return
		}

		//Flushing the writer buffer
		err = writer.Flush()
		if err != nil {
			fmt.Println("Failed to flush writer")
			return
		}

	}
}

func StartServer(HOST string, PORT string, TYPE string) {
	for {
		server, err := net.Listen(TYPE, HOST+":"+PORT)
		if err != nil {
			fmt.Println("failed to start the server: ", err.Error())
			continue
		}
		fmt.Println("Server Started")
		for {
			//accept incoming connections
			conn, err := server.Accept()
			if err != nil {
				fmt.Println("Failed to accept connection")
				break
			}
			fmt.Println("New connection: ", conn.RemoteAddr().String())

			//handle the connection in a separate go routine
			go handleConnection(conn)
		}

		if server.Close() != nil {
			fmt.Println("Failed to close the server")
		} else {
			fmt.Print("Server closed successfully")
		}
	}
}

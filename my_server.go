package main

import (
	"bufio"
	"bytes"
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

	reader, writer := bufio.NewReader(bytes.NewBufferString("hello world$")), bufio.NewWriter(conn)
	dupeReader := bufio.NewReader(conn)
	for {
		//Read data from the connection
		// Handle the data reading part somewhere else I guess ????

		reader = bufio.NewReader(bytes.NewBufferString("hello world$"))

		data, err := reader.ReadBytes('$')
		HandleRequest(dupeReader)
		//data, err := bufio.NewReader(bytes.NewBufferString("sus amongus\n")).ReadBytes('\n')

		if err != nil {
			fmt.Println("Failed to read data: ", err.Error())
			return
		}
		fmt.Println("Data read successfully", string(data[:]))
		// check first 3 characters of the message to know if it is a get, set, put, del???

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

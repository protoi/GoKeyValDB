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

	//reader := bufio.NewReader(bytes.NewBufferString("hello world$"))
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	//making a hashmap

	db := make(map[string]string)

	for {

		s, b, i := HandleRequest(reader, &db)

		ack := fmt.Sprintf("=> %v %v %v", s, b, i)

		_, err := writer.WriteString(ack)
		if err != nil {
			fmt.Println("Failed to write data: ", err.Error())
			return
		}

		//Flushing the writer buffer
		if writer.Size() > 0 {
			if err = writer.Flush(); err != nil {
				fmt.Println("Failed to flush writer")
				return
			}
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

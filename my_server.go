package main

import (
	"bufio"
	"fmt"
	"net"
)

type DataStructureCollection struct {
	kv_data *map[string]*KeyValMapping
	ll_data *map[string]*BiDirectionalLinkedList
	sl_data *map[string]*SkipList
}

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

	//making a hashmap of string string pai
	//db := make(map[string]string)

	// vanilla key val map + linked list map, skip list map
	kv_ds := make(map[string]*KeyValMapping)
	ll_ds := make(map[string]*BiDirectionalLinkedList)
	sl_ds := make(map[string]*SkipList)

	user := DataStructureCollection{
		kv_data: &kv_ds,
		ll_data: &ll_ds,
		sl_data: &sl_ds,
	}

	for {

		s, b, i := HandleRequest(reader, &user)

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

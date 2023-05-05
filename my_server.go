package main

import (
	"bufio"
	"fmt"
	"net"
	"regexp"
)

type DataStructureCollection struct {
	kv_data *map[string]*KeyValMapping
	ll_data *map[string]*BiDirectionalLinkedList
	sl_data *map[string]*SkipList
}
type UserInformation struct {
	userID             string
	authToken          string
	userDataStructures *DataStructureCollection
}

func handleConnection(conn net.Conn, allUsers *map[string]*UserInformation) {
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

	userInfo := UserInformation{}
	// make a registration call?
	// the first call has to be a username and access token setting call

	//var userID string

	if dataRead, success := ReadBuffer(reader); success {

		fmt.Println("Username And Access Token = ", dataRead)
		var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(.*)$`)
		match := re.FindStringSubmatch(dataRead)
		if len(match) == 3 {
			userID := match[1]
			authToken := match[2]

			// if this username already exists then halt
			if _, ok := (*allUsers)[userID]; ok == false { // this user does not exist
				userInfo.userID = userID
				userInfo.authToken = authToken
				(*allUsers)[userID] = &userInfo
				fmt.Println("Logger in as " + userID + " with token " + authToken)

				// send initiation ack to user
				_, err := writer.WriteString("Logger in as " + userID + " with token " + authToken)
				if err != nil {
					fmt.Println("Failed to create user: ", err.Error())
					return
				}

				//Flushing the writer buffer
				if writer.Size() > 0 {
					if err = writer.Flush(); err != nil {
						fmt.Println("Failed to flush writer")
						return
					}
				}
			} else { // user exists already
				// new user trying to log in with the same userID
				// break out of this after sending a message saying this is not allowed

				// send ack to user
				_, err := writer.WriteString("overwriting authTokens are not allowed")
				if err != nil {
					fmt.Println("Failed to login: ", err.Error())
					return
				}

				//Flushing the writer buffer
				if writer.Size() > 0 {
					if err = writer.Flush(); err != nil {
						fmt.Println("Failed to flush writer")
						return
					}
				}
				return // drop the entire connection
			}
		}
	}

	//making a hashmap of string string pai
	//db := make(map[string]string)

	// vanilla key val map + linked list map, skip list map
	kv_ds := make(map[string]*KeyValMapping)
	ll_ds := make(map[string]*BiDirectionalLinkedList)
	sl_ds := make(map[string]*SkipList)

	userData := DataStructureCollection{
		kv_data: &kv_ds,
		ll_data: &ll_ds,
		sl_data: &sl_ds,
	}

	userInfo.userDataStructures = &userData

	for {
		// string, bool, int
		s, b, i := HandleRequest(reader, allUsers, &userInfo)

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
	allUsers := make(map[string]*UserInformation)

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
			go handleConnection(conn, &allUsers)
		}

		if server.Close() != nil {
			fmt.Println("Failed to close the server")
		} else {
			fmt.Print("Server closed successfully")
		}
	}
}

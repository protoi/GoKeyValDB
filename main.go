package main

import (
	"fmt"
)

const (
	HOST = "localhost"
	PORT = "4000"
	TYPE = "tcp"
)

func main() {
	fmt.Println("hello world")
	StartServer(HOST, PORT, TYPE)
	//runner()
}

/*
TODO: allow for cross user interaction, let one user interact with another users data structures.
	Obviously have some sort of authentication though.
	Maybe a token issued when a hashset, linkedlist or zset is created
*/

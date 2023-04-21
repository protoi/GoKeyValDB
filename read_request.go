package main

import (
	"bufio"
	"fmt"
)

func HandleRequest(reader *bufio.Reader) {
	ans, success := ReadBuffer(reader)
	if !success {
		return
	}
	fmt.Println("---> ", ans)
}

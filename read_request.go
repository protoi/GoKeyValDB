package main

import (
	"bufio"
	"fmt"
)

func HandleRequest(reader *bufio.Reader, db *map[string]string) (string, bool, int) {
	if dataRead, success := ReadBuffer(reader); success {
		fmt.Println("---> ", dataRead)

		s, b, i := PerformAction(db, dataRead)

		fmt.Printf("%v %v %v\n", s, b, i)

		return s, b, i
	}
	return "", false, 0
}

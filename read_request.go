package main

import (
	"bufio"
	"fmt"
)

func HandleRequest(reader *bufio.Reader, allUsers *map[string]*UserInformation, userInfo *UserInformation) (string, bool, int) {
	if dataRead, success := ReadBuffer(reader); success {
		fmt.Println("---> ", dataRead)

		s, b, i := PerformAction(dataRead, allUsers, userInfo)

		fmt.Printf("%v %v %v\n", s, b, i)

		return s, b, i
	}
	return "", false, 0
}

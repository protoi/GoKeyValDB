package main

import (
	"bufio"
	"strconv"
)

func getActualMessageLen(reader **bufio.Reader) (int, bool) {
	data, err := (*reader).ReadBytes('%')
	if err != nil {
		return -1, false
	}
	dataWithDelimRemoved := data[:len(data)-1]
	actualMessageLength, err := strconv.Atoi(string(dataWithDelimRemoved[:]))
	if err != nil {
		return -1, false
	}
	return actualMessageLength, true
}

func ReadBuffer(reader *bufio.Reader) (string, bool) {

	msgLen, success := getActualMessageLen(&reader)

	// if unsuccessful
	if success != true {
		return "", false
	}

	bufferContent := make([]byte, msgLen)

	bytesRead, err := reader.Read(bufferContent)
	if err != nil {
		return "", false
	}

	actualMessage := string(bufferContent[:bytesRead])
	return actualMessage, true
}

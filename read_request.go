package main

import (
	"bufio"
	"fmt"
	"io"
)

type clientInformation struct {
	command    *string // can be either strings or nil
	identifier *string
	value      *string
}

// region parsing information

func parseCommand(reader *bufio.Reader) (*string, error) {
	parsedCommand, err := reader.ReadBytes(' ')
	if err != nil && err != io.EOF {
		return nil, err
	}
	fmt.Println("parsedCommand --> ", string(parsedCommand[:]))
	strParsedCommand := string(parsedCommand[:len(parsedCommand)-1])
	return &strParsedCommand, nil
}
func parseID(reader *bufio.Reader) (*string, error) {
	parsedID, err := reader.ReadBytes(' ')
	if err != nil && err != io.EOF {
		return nil, err
	}
	fmt.Println("parsedID --> ", string(parsedID[:]))
	strParsedID := string(parsedID[:len(parsedID)-1])
	return &strParsedID, nil
}
func parseValue(reader *bufio.Reader) (*string, error) {
	parsedValue, err := reader.ReadBytes(' ')
	if err != nil && err != io.EOF {
		return nil, err
	}
	fmt.Println("parsedValue --> ", string(parsedValue[:]))
	strParsedValue := string(parsedValue[:len(parsedValue)-1])
	return &strParsedValue, nil
}

//endregion

/*
SET key value - sets the value of a key in Redis to a given string value. ✅
GET key - retrieves the value of a key in Redis. ✅
DEL key - deletes a key and its value from Redis. ✅
INCR key - increments the value of a key by 1. ✅
DECR key - decrements the value of a key by 1. ✅
EXPIRE key seconds - sets a time-to-live (TTL) for a key in Redis, after which the key and its value will be deleted.
KEYS pattern - returns all keys in Redis that match a given pattern.
FLUSHDB - deletes all keys and their values from the current Redis database. ✅
PING - checks if the Redis server is running and responds with "PONG".
*/
func handleCases(reader *bufio.Reader) *clientInformation {

	// TODO: command ID and value behaving weirdly, put breakpoints and check their values. information retrieved "cmd: set, val: fkdj, id: set" for "set fkdj "

	var tempCmd, tempId, tempVal *string = nil, nil, nil

	// checking commands
	cmd, err := parseCommand(reader)
	if err != nil || cmd == nil { // command read unsuccessful
		return nil
	}
	switch *cmd {
	case "flush":
		tempCmd = cmd

	case "set": // has a key and a value associated with it
		id, err := parseID(reader)
		if err != nil || id == nil {
			return nil // key read unsuccessful
		}

		val, err := parseValue(reader)
		if err != nil || val == nil {
			return nil // value read unsuccessful
		}
		tempCmd, tempId, tempVal = cmd, id, val

	case "get", "del", "incr", "decr": // only key associated with this command
		id, err := parseID(reader)
		if err != nil || id == nil {
			return nil
		}
		tempCmd, tempId = cmd, id
	default:
		return nil
	}

	info := clientInformation{
		command:    tempCmd,
		identifier: tempId,
		value:      tempVal,
	}
	return &info
}

func HandleRequest(reader *bufio.Reader) {
	// takes in a reader and reads the buffer contents
	// assume delimiter is a ' ' sign

	info := handleCases(reader)
	if info != nil {
		cmd, id, val := (*info).command, (*info).identifier, (*info).value
		var cmd_, id_, val_ string

		if cmd != nil {
			cmd_ = *cmd
		}
		if id != nil {
			id_ = *id
		}
		if val != nil {
			val_ = *val
		}
		fmt.Printf("information retreived cmd: %s, val: %s, id: %s\n", cmd_, id_, val_)
		return
	}
	fmt.Println("oh no something went wrong")
}

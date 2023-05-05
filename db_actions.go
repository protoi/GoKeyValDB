package main

import (
	"fmt"
	"strings"
)

/*
SET key value - sets the value of a key in Redis to a given string value. ✅
GET key - retrieves the value of a key in Redis. ✅
DEL key - deletes a key and its value from Redis. ✅
INCR key - increments the value of a key by 1. ✅
DECR key - decrements the value of a key by 1. ✅
FLUSHDB - deletes all keys and their values from the current Redis database. ✅


	TODO: try adding LISTPUSH and LISTPOP and LISTRANGE -> https://www.tutorialspoint.com/redis/redis_lists.htm
	TODO: SETADD, SETPOP, SETMEMBERS, SETDIFF -> https://www.tutorialspoint.com/redis/redis_sets.htm
	TODO: ZSETADD, ZSETPOP


*/

func PerformAction(readData string, allUsers *map[string]*UserInformation, userInfo *UserInformation) (string, bool, int) {

	// splitting upon a space
	command, substance := "", ""
	if indexOfCommand := strings.Index(readData, " "); indexOfCommand > -1 {
		command = readData[:indexOfCommand]
		substance = readData[indexOfCommand+1:]
	} else {
		command = readData
	}

	var userID = userInfo.userID
	user := (*allUsers)[userID].userDataStructures
	switch command {
	case "LIST":
		element, status := HandleLinkedList(substance, user.ll_data)
		return element, status, 0
	case "KV":
		element, status := HandleKeyValMapping(substance, user.kv_data)
		return element, status, 0

	case "ZSET":
		key, value, status := HandleSkipList(substance, user.sl_data)
		return key, status, value

	default:
		// invalid commands
		fmt.Println("invalid commands")
	}

	return "", false, 0
}

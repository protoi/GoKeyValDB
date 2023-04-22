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
EXPIRE key seconds - sets a time-to-live (TTL) for a key in Redis, after which the key and its value will be deleted.
KEYS pattern - returns all keys in Redis that match a given pattern.
FLUSHDB - deletes all keys and their values from the current Redis database. ✅
PING - checks if the Redis server is running and responds with "PONG".
*/

func setValue(key string, value string, db *map[string]string) bool {
	if db == nil {
		return false
	}
	(*db)[key] = value
	fmt.Printf("db[%s] <- %s\n", key, value)
	return true
}

func getValue(key string, db *map[string]string) (string, bool) {
	if db == nil {
		return "", false
	}
	val, status := (*db)[key]
	if status {
		fmt.Printf("db[%s] -> %s\n", key, val)
	}
	return val, status
}

func deleteValue(key string, db *map[string]string) bool {
	if db == nil {
		return false
	}
	if _, status := (*db)[key]; status {
		// key present so you can delete it
		delete(*db, key)
		fmt.Printf("db[%s] deleted\n", key)
		return true
	}
	return false
}

func flushDB(db *map[string]string) bool {
	if db == nil {
		return false
	}
	for key := range *db {
		delete(*db, key)
	}
	fmt.Println("db flushed")
	return true
}

func dbLength(db *map[string]string) (int, bool) {
	if db == nil {
		return -1, false
	}
	fmt.Printf("len -> %d\n", len(*db))
	return len(*db), true
}

func extractKeyVal(data string) (string, string, bool) {
	if indexOfKey := strings.Index(data, " "); indexOfKey > -1 {
		key, val := data[:indexOfKey], data[indexOfKey+1:]
		return key, val, true
	}
	return "", "", false
}

func PerformAction(db *map[string]string, readData string) (string, bool, int) {

	// splitting upon a space
	command, substance := "", ""
	if indexOfCommand := strings.Index(readData, " "); indexOfCommand > -1 {
		command = readData[:indexOfCommand]
		substance = readData[indexOfCommand+1:]
	} else {
		command = readData
	}

	key, val, status := "", "", false
	switch command {
	case "set":
		keyValStatus := false
		if key, val, keyValStatus = extractKeyVal(substance); keyValStatus {
			status = setValue(key, val, db)
		}
		return "", status, 0
	case "get":
		key = substance
		val, status = getValue(key, db)
		return val, status, 0
	case "del":
		key = substance
		status = deleteValue(key, db)

		return "", status, 0
	case "flush":
		status = flushDB(db)
		return "", status, 0
	case "len":
		length := -1
		length, status = dbLength(db)
		return "", status, length
	default:
		// invalid commands
		fmt.Println("invalid commands")
	}

	return "", false, 0
}

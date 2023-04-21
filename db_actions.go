package main

import "fmt"

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

func SetValue(key string, value string, db *map[string]string) bool {
	if db == nil {
		return false
	}
	(*db)[key] = value
	fmt.Printf("db[%s] <- %s\n", key, value)
	return true
}

func GetValue(key string, db *map[string]string) (string, bool) {
	if db == nil {
		return "", false
	}
	val, status := (*db)[key]
	if status {
		fmt.Printf("db[%s] -> %s\n", key, val)
	}
	return val, status
}

func DeleteValue(key string, db *map[string]string) bool {
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

func FlushDB(db *map[string]string) bool {
	if db == nil {
		return false
	}
	for key := range *db {
		delete(*db, key)
	}
	fmt.Println("db flushed")
	return true
}

func DBLength(db *map[string]string) (int, bool) {
	if db == nil {
		return -1, false
	}
	fmt.Printf("len -> %d\n", len(*db))
	return len(*db), true
}

func Decide(db *map[string]string, readData string) {

}

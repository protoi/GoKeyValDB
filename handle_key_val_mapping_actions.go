package main

import (
	"fmt"
	"regexp"
)

// HandleKeyValMapping takes the command issued and the performs action with hashmap dictionary provided
func HandleKeyValMapping(substance string, KVMapping *map[string]*KeyValMapping) (string, bool) {
	/*
		KV has already been consumed

		make
		1. kv init <map_name>
		2. kv set <mapname> <key> <val>
		3. kv get <mapname> <key>
		4. kv del <mapname> <key>
		5. kv flush <mapname>
	*/
	command, restOfTheString, status := KVgetCommandAndRest(substance)

	if status == false {
		return "", false
	}
	switch command {
	case "INIT":
		if mapName, status := KVgetMapNameOnly(restOfTheString); status {
			if _, ok := (*KVMapping)[mapName]; !ok {
				tempKVMap := &KeyValMapping{}
				tempKVMap.init()
				(*KVMapping)[mapName] = tempKVMap
				return "", true
			}
		}
	case "GET":
		if mapName, key, status := KVgetMapNameAndKey(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return kvmap.GetValue(key)
			}
		}
	case "SET":
		if mapName, key, val, status := KVgetMapNameKeyAndValue(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return "", kvmap.SetValue(key, val)
			}
		}
	case "DEL":
		if mapName, key, status := KVgetMapNameAndKey(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return "", kvmap.DelValue(key)
			}
		}
	case "FLUSH":
		if mapName, status := KVgetMapNameOnly(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return "", kvmap.FlushMap()
			}
		}
	default:
		fmt.Println("Invalid command detected: " + substance)
	}

	return "", false
}

// KVgetCommandAndRest takes a string and returns the command and the rest of the string
func KVgetCommandAndRest(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)(.*)$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}

// KVgetMapNameKeyAndValue takes a string and returns the name of hashmap, the key and the value
func KVgetMapNameKeyAndValue(substance string) (string, string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s+(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 4 {
		return match[1], match[2], match[3], true
	}
	return "", "", "", false
}

// KVgetMapNameAndKey takes a string and returns the name of hashmap and the key
func KVgetMapNameAndKey(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}

// KVgetMapNameOnly takes a string and returns the name of the hashmap
func KVgetMapNameOnly(substance string) (string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

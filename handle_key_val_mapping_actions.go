package main

import (
	"fmt"
	"regexp"
)

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
		if mapName, status := KVgetMapnameOnly(restOfTheString); status {
			if _, ok := (*KVMapping)[mapName]; !ok {
				tempKVMap := &KeyValMapping{}
				tempKVMap.init()
				(*KVMapping)[mapName] = tempKVMap
				return "", true
			}
		}
	case "GET":
		if mapName, key, status := KVgetMapnameAndKey(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return kvmap.GetValue(key)
			}
		}
	case "SET":
		if mapName, key, val, status := KVgetMapnameKeyAndValue(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return "", kvmap.SetValue(key, val)
			}
		}
	case "DEL":
		if mapName, key, status := KVgetMapnameAndKey(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return "", kvmap.DelValue(key)
			}
		}
	case "FLUSH":
		if mapName, status := KVgetMapnameOnly(restOfTheString); status {
			if kvmap, ok := (*KVMapping)[mapName]; ok {
				return "", kvmap.FlushMap()
			}
		}
	default:
		fmt.Println("Invalid command detected: " + substance)
	}

	return "", false
}

func KVgetCommandAndRest(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)(.*)$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}

func KVgetMapnameKeyAndValue(substance string) (string, string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s+(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 4 {
		return match[1], match[2], match[3], true
	}
	return "", "", "", false
}

func KVgetMapnameAndKey(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}

func KVgetMapnameOnly(substance string) (string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

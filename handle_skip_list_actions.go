package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func HandleSkipList(substance string, skipListMapping *map[string]*SkipList) (string, int, bool) {
	/*
		input like ->
		zset init <skiplist_name>
		zset SearchKey <skiplist_name> <string>
		zset SearchVal <skiplist_name> <int>
		zset Insert <skiplist_name> <string> <int>
		zset Delete <skiplist_name> <string>
		zset GetMax <skiplist_name>
		zset PopMax <skiplist_name>
		zset GetMin <skiplist_name>
		zset PopMin <skiplist_name>
		zset GetPredecessor <skiplist_name> <int>
		zset GetSuccessor <skiplist_name> <int>
		zset PrintList <skiplist_name>
	*/

	command, restOfTheString, status := SLgetCommandAndRest(substance)

	if status == false {
		return "", -1, false
	}

	switch command {
	case "INIT":
		if skipListName, status := SLgetSkipListNameOnly(restOfTheString); status {
			// check if skip list of this name does not exist yet
			if _, ok := (*skipListMapping)[skipListName]; ok == false {
				tempSkipList := &SkipList{}
				tempSkipList.init(10000)
				(*skipListMapping)[skipListName] = tempSkipList
				return "", -1, true
			}
		}
	case "SEARCHKEY":
		if skipListName, key, status := SLgetSkipListNameAndKey(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				value, found := sl.SearchKey(key)
				return "", value, found
			}
		}
	case "SEARCHVAL":
		if skipListName, value, status := SLgetSkipListNameAndValue(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return "", -1, sl.SearchVal(value)
			}
		}
	case "INSERT":
		if skipListName, keyToInsert, ValueToInsert, status := SLgetSkipListNameKeyAndValue(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				sl.Insert(keyToInsert, ValueToInsert)
				return "", -1, true
			}
		}
	case "DELETE":
		if skipListName, key, status := SLgetSkipListNameAndKey(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return "", -1, sl.Delete(key)
			}
		}
	case "GETMAX":
		if skipListName, status := SLgetSkipListNameOnly(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return sl.GetMax()
			}
		}
	case "GETMIN":
		if skipListName, status := SLgetSkipListNameOnly(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return sl.GetMin()
			}
		}
	case "POPMAX":
		if skipListName, status := SLgetSkipListNameOnly(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return sl.PopMax()
			}
		}
	case "POPMIN":
		if skipListName, status := SLgetSkipListNameOnly(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				sl.PrintList()
				return sl.PopMin()
			}
		}
	case "GETPREV":
		if skipListName, value, status := SLgetSkipListNameAndValue(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return sl.GetPredecessor(value)
			}
		}
	case "GETNEXT":
		if skipListName, value, status := SLgetSkipListNameAndValue(restOfTheString); status {
			if sl, ok := (*skipListMapping)[skipListName]; ok {
				return sl.GetSuccessor(value)
			}
		}
	default:
		fmt.Println("Invalid command detected " + substance)
	}
	return "", -1, false

}

func SLgetCommandAndRest(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)(.*)$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}

func SLgetSkipListNameAndKey(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}

	return "", "", false
}

func SLgetSkipListNameAndValue(substance string) (string, int, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\d+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		if integerValue, err := strconv.Atoi(match[2]); err == nil {
			return match[1], integerValue, true
		}
	}
	return "", -1, false
}

func SLgetSkipListNameKeyAndValue(substance string) (string, string, int, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s+(\d+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 4 {
		if integerValue, err := strconv.Atoi(match[3]); err == nil {
			return match[1], match[2], integerValue, true
		}
	}
	return "", "", -1, false
}

func SLgetSkipListNameOnly(substance string) (string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false
}

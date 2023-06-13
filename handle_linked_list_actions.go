package main

import (
	"fmt"
	"regexp"
)

func HandleLinkedList(substance string, linkedListMapping *map[string]*BiDirectionalLinkedList) (string, bool) {
	// TODO: add LISTPUSH and LISTPOP and LISTRANGE -> https://www.tutorialspoint.com/redis/redis_lists.htm

	/*
		"list" has been consumed already
		make
		1. list init <list_name> ✅
		2. list push_front <list_name> <element> ✅
		3. list pop_front <list_name> <element> ✅
		4. list push_back <list_name> ✅
		5. list pop_back <list_name> ✅
		6. list peek_front <list_name> ✅
		7. list peek_back <list_name> ✅
	*/

	command, restOfTheString, status := LLgetCommandAndRest(substance)
	if status == false {
		return "", false
	}
	switch command {
	case "INIT":
		if listName, status := LLgetListNameOnly(restOfTheString); status == true {
			// list of this name does not exist yet
			if _, ok := (*linkedListMapping)[listName]; ok == false {
				(*linkedListMapping)[listName] = &BiDirectionalLinkedList{}
				return "", true
			}
		}

	case "PUSHFRONT":
		if listName, elementToBePushed, status := LLgetListNameAndElement(restOfTheString); status == true {
			// check if linked list of this name is present for this user
			if ll, ok := (*linkedListMapping)[listName]; ok {
				ll.PushBack(elementToBePushed)
				return "", true
			}
		}
	case "PUSHBACK":
		if listName, elementToBePushed, status := LLgetListNameAndElement(restOfTheString); status == true {
			// check if linked list of this name is present for this user
			if ll, ok := (*linkedListMapping)[listName]; ok {
				ll.PushFront(elementToBePushed)
				return "", true
			}
		}
	case "POPFRONT":
		if listName, status := LLgetListNameOnly(restOfTheString); status == true {
			// check if linked list of this name is present for this user
			if ll, ok := (*linkedListMapping)[listName]; ok {
				return ll.PopFront()
			}
		}
	case "POPBACK":
		if listName, status := LLgetListNameOnly(restOfTheString); status == true {
			// check if linked list of this name is present for this user
			if ll, ok := (*linkedListMapping)[listName]; ok {
				return ll.PopBack()
			}
		}
	case "PEEKFRONT":
		if listName, status := LLgetListNameOnly(restOfTheString); status == true {
			// check if linked list of this name is present for this user
			if ll, ok := (*linkedListMapping)[listName]; ok {
				return ll.PeekFront()
			}
		}
	case "PEEKBACK":
		if listName, status := LLgetListNameOnly(restOfTheString); status == true {
			// check if linked list of this name is present for this user
			if ll, ok := (*linkedListMapping)[listName]; ok {
				return ll.PeekBack()
			}
		}
	default:
		fmt.Println("Invalid command detected: " + substance)
	}
	return "", false
}

/*
input -> "push_back mylist aaaaaaa"
output -> "push_back", " mylist aaaaaaa"
*/
func LLgetCommandAndRest(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)(.*)$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}
func LLgetListNameAndElement(substance string) (string, string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s+(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 3 {
		return match[1], match[2], true
	}
	return "", "", false
}
func LLgetListNameOnly(substance string) (string, bool) {
	var re = regexp.MustCompile(`(?m)^\s*(\w+)\s*$`)
	match := re.FindStringSubmatch(substance)
	if len(match) == 2 {
		return match[1], true
	}
	return "", false

}

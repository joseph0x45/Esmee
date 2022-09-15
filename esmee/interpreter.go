package main

import (
	"strings"
)

func interpreter(command string) (result string) {
	switch  {
	case command=="Help":
		result = "Welcome to EsmeeDB exposed CLI"
	case command== "help":
		result = "Welcome to EsmeeDB exposed CLI"
	case command== "exit":
		result = "Bye"
	case strings.Contains(command, "create") :
		treeName:= strings.Split(command, " ")
		result = treeName[0]
	}

	return
	
}
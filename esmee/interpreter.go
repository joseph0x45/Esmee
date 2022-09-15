package main

import (
	"fmt"
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
		subCommands := strings.Split(command, " ")
		fmt.Println(len(subCommands))
		switch  {
		case len(subCommands)<5:
			fmt.Println("Invalid syntax")
			result = "Invalid syntax"
		}
	}

	return
	
}
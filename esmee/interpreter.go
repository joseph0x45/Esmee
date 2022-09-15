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
		case strings.Contains(command, "climb"):
			subCommands := strings.Split(command, " ")
			switch{
				case len(subCommands)>2 || len(subCommands)<2:
					result = "Invalid syntax. Climb requires only one argument which is the tree name"
				case len(subCommands)==2:
					treeName := subCommands[1]
					result = treeName
			}
		case strings.Contains(command, "create") :
			subCommands := strings.Split(command, " ")
			switch  {
				case len(subCommands)>3 || len(subCommands)<3 :
					result = "Invalid syntax"
				case subCommands[1]=="tree":
					result = fmt.Sprintf("wants to create a tree named %v", subCommands[2])
				case subCommands[1]=="branch":
					result = fmt.Sprintf("wants to create a branch named %v", subCommands[2])
			}
		default:
			fmt.Printf("<%v>", command)
	}

	return
	
}
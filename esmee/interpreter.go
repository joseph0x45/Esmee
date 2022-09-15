package main

import (
	"fmt"
	"strings"
)

var rootTree string = "root/"

func pathResolver(path string) (resolvedPath string) {
	resolvedPath = fmt.Sprintf("%v%v", rootTree, path)
	return
}

func interpreter(command string) (result string) {
	switch {

	//ANCHOR Help and exit
	case command == "Help":
		result = "Welcome to EsmeeDB exposed CLI"
	case command == "help":
		result = "Welcome to EsmeeDB exposed CLI"
	case command == "exit":
		result = "Bye"

	//ANCHOR CLIMB
	case strings.Contains(command, "climb"):
		subCommands := strings.Split(command, " ")
		switch {
		case len(subCommands) > 2 || len(subCommands) < 2:
			result = "Invalid syntax. Climb requires only one argument which is the tree name"
		case len(subCommands) == 2:
			treeName := subCommands[1]
			switch {
				case treeName=="down" :
					if rootTree!="root/" {
						treeName:= strings.Split(rootTree, "/")
						result = fmt.Sprintf("climbdown %v", treeName[1])
						rootTree = "root/"
						return
					}
					rootTree = fmt.Sprintf("%v%v", rootTree, treeName)
					result = fmt.Sprintf("climb %v", treeName)
				default:
					rootTree = fmt.Sprintf("%v%v", rootTree, treeName)
					result = fmt.Sprintf("climb %v", treeName)
			}

		}

		//ANCHOR CREATES
	case strings.Contains(command, "create"):
		subCommands := strings.Split(command, " ")
		switch {
		case len(subCommands) > 3 || len(subCommands) < 3:
			result = "Invalid syntax"
		case subCommands[1] == "tree":
			result = fmt.Sprintf("wants to create a tree named %v", subCommands[2])
		case subCommands[1] == "branch":
			result = fmt.Sprintf("wants to create a branch named %v", subCommands[2])
		}
	default:
		fmt.Printf("Unrecognized command <%v>", command)
	}

	return

}

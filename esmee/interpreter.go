package main

import (
	"fmt"
	"strings"
)

const helpMessage string = 
`Welcome to EsmeeDB CLI
This is a short documentation of the commands available. Refeer to the online docs for more details
`

var rootTree string = "root/"

func pathResolver(path string) (resolvedPath string) {
	resolvedPath = fmt.Sprintf("%v%v", rootTree, path)
	return
}

func interpreter(command string) (result string) {
	switch {

	//ANCHOR Help and exit
	case command == "Help":
		fmt.Print(helpMessage)
	case command == "help":
		fmt.Print(helpMessage)
	case command == "exit":
		fmt.Println("Bye")
		result = "Bye"

	//ANCHOR CLIMB & CLIMBDOWN
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
					rootTree = fmt.Sprintf("%v%v/", rootTree, treeName)
					result = fmt.Sprintf("climb %v", treeName)
				default:
					rootTree = fmt.Sprintf("%v%v/", rootTree, treeName)
					result = fmt.Sprintf("climb %v", treeName)
			}

		}

		//ANCHOR CREATES
	case strings.Contains(command, "create tree"):
		subCommands := strings.Split(command, " ")
		switch {
			case len(subCommands) > 3 || len(subCommands) < 3:
				result = "Error IS"
			case rootTree!="root/":
				result = "Error TIT"
			default :
				treePath := pathResolver(subCommands[2])
				_, result = createTree(treePath)
		}
	

	case strings.Contains(command, "create branch"):
		subCommands := strings.Split(command, " ")
		switch  {
			case len(subCommands) >3 || len(subCommands)<3:
				result = "Error IS"
			case rootTree=="root/":
				result = "Error BOOT"
			default :
				branchName := subCommands[2]
				_, result = createBranch(pathResolver(branchName))
		}

	default:
		result = "Error UC"
	}

	return

}

package main

import (
	"fmt"
	"strings"
	"os"
)

const helpMessage string = `Welcome to EsmeeDB CLI
This is a short documentation of the commands available. Refeer to the online docs for more details
>create tree/branch:
    The create commmand is used to create a new tree or branch. To create a branch you must be in a tree
Lets say we created a tree called Food (using create tree Food). In order to create a branch "Vegetarian" on the tree
we will first have to climb up the tree.
>climb Food
>Food#
once we climb up a tree, the prompt is updated with the name of the tree. And only then we can create a branch
>Food#create branch Vegetarians
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
			case treeName == "down":
				if rootTree != "root/" {
					treeName := strings.Split(rootTree, "/")
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
		case rootTree != "root/":
			result = "Error TIT"
		default:
			treePath := pathResolver(subCommands[2])
			_, result = createTree(treePath)
		}

	case strings.Contains(command, "create branch"):
		subCommands := strings.Split(command, " ")
		switch {
		case len(subCommands) > 3 || len(subCommands) < 3:
			result = "Error IS"
		case rootTree == "root/":
			result = "Error BOOT"
		default:
			branchName := subCommands[2]
			_, result = createBranch(pathResolver(branchName))
		}

	
	//ANCHOR Get forest arborescence
	case command=="arb":
		content, err := os.ReadDir(rootTree)
		if err!= nil{
			fmt.Println(err)
		}
		for _, file := range content{
			nature := ""
			switch {
			case file.IsDir() && rootTree=="root/" :
				nature = "🌴"
			case file.IsDir() :
				nature = "🌿"
			default:
				nature = "🍀"
			}
			fmt.Printf("%v%v%v\n",nature, file.Name(), nature)
		}
	
	default:
		result = "Error UC"
	}

	return

}

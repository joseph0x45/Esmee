package esmee

import (
	"errors"
	"fmt"
	"strings"
)

const helpMessage string = `
Welcome to EsmeeDB CLI

This is a short documentation of the commands available. Refer to the online docs [TODO] for more details	

$ create [tree/branch]
The create commmand is used to create a new tree or branch

$ create tree
We can use this command to create a tree - $ create tree Food

$ climb [name of tree]
This is used to enter/climb a tree, e.g $ climb Vegetarian.
Once we climb up a tree, the prompt is updated with the name of the tree. And only then we can 
create a branch. e.g $ Food#

$ create branch:
To create a branch you must be in a tree. In order to create a branch "Vegetarian" on the tree
Inside the tree, we do # create branch Vegetarian
	
>Food#create branch Vegetarians
`

type Command struct {
	mainCommand string
	subCommand  []string
}

var rootTree string = "root/"

func pathResolver(path string) (resolvedPath string) {
	resolvedPath = fmt.Sprintf("%v%v", rootTree, path)
	return
}

func sanitizeCommand(input string) (Command, error) {
	input = strings.TrimSpace(input)

	if len(input) == 0 {
		return Command{}, errors.New("please provide a command")
	}

	// climb tree becomes [climb tree]
	commandsBlueprint := strings.SplitN(input, " ", 2)

	command := new(Command)

	command.mainCommand = commandsBlueprint[0]
	if len(commandsBlueprint) > 1 {
		command.subCommand = strings.Split(commandsBlueprint[1], " ")
	} else {
		command.subCommand = []string{}
	}

	return *command, nil
}

func interpreter(cmd string) (string, error) {
	result := ""
	command, err := sanitizeCommand(cmd)

	if err != nil {
		return "", err
	}

	fmt.Printf("Command: %+v", command)

	// Once user command is retrieved, parse it
	switch {
	//ANCHOR Help and exit
	case strings.ToLower(command.mainCommand) == "help":
		fmt.Print(helpMessage)
	case strings.ToLower(command.mainCommand) == "exit":
		fmt.Println("Bye")
		result = "Bye"

	// ANCHOR CREATES
	case command.mainCommand == "create":

		// ANCHOR TREE
		if command.subCommand[0] == "tree" {
			switch {
			case len(command.subCommand) != 2:
				return "", errors.New("invalid argument length. the create command takes two arguments")
			case rootTree != "root/":
				return "", errors.New("a tree cannot be created in another tree...thats weird")
			default:
				treePath := pathResolver(command.subCommand[1])
				fmt.Println("The tree path is: ", treePath)
				switch treePath {
				case fmt.Sprintf("%vdown", rootTree):
					result = "Error CDRK"
				default:
					_, result = createTree(treePath)
				}
			}
		}

		// //ANCHOR CLIMB & CLIMBDOWN
		// case command.mainCommand == "climb":
		// 	switch {
		// 	case len(command.subCommand) > 2 || len(command.subCommand) < 2:
		// 		return "", errors.New("climb command requires only one argument - $ climb treename")
		// 	case len(command.subCommand) == 2:
		// 		treeName := command.subCommand[0]
		// 		switch {
		// 		case treeName == "down":
		// 			if rootTree != "root/" {
		// 				treeName := strings.Split(rootTree, "/")
		// 				result = fmt.Sprintf("climbdown %v", treeName[1])
		// 				rootTree = "root/"
		// 				return result, nil
		// 			} else {
		// 				fmt.Println("You can't use `down` when not in a tree.")
		// 				return result, nil
		// 			}
		// 		default:
		// 			rootTree = fmt.Sprintf("%v%v/", rootTree, treeName)
		// 			_, err := os.ReadDir(rootTree)
		// 			if err != nil {
		// 				result = "Error ITOB"
		// 				rootTree = "root/"
		// 			} else {
		// 				result = fmt.Sprintf("climb %v", treeName)
		// 			}
		// 		}

		// 	}

		// case strings.Contains(command, "create branch"):
		// 	subCommands := strings.Split(command, " ")
		// 	switch {
		// 	case len(subCommands) > 3 || len(subCommands) < 3:
		// 		result = "Error IS"
		// 	case rootTree == "root/":
		// 		result = "Error BOOT"
		// 	default:
		// 		branchName := subCommands[2]
		// 		fmt.Println(branchName)
		// 		switch branchName {
		// 		case "down":
		// 			result = "Error CDRK"
		// 		default:
		// 			_, result = createBranch(pathResolver(branchName))
		// 		}
		// 	}

		// //ANCHOR Get forest arborescence
		// case command == "arb":
		// 	content, err := os.ReadDir(rootTree)
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	}
		// 	for _, file := range content {
		// 		nature := ""
		// 		switch {
		// 		case file.IsDir() && rootTree == "root/":
		// 			nature = "🌴"
		// 		case file.IsDir():
		// 			nature = "🌿"
		// 		default:
		// 			nature = "🍀"
		// 		}
		// 		fmt.Printf("%v%v%v\n", nature, file.Name(), nature)
		// 	}

		// default:
		// 	result = "Error UC"
	}

	return result, nil

}

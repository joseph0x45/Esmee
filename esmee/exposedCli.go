package main

import (
	"fmt"
)

func commandBuilder(commandArray [3]string) (command string) {
	if commandArray[1]=="" || commandArray[2]=="" {
		command = commandArray[0]
		return
	}
	for _, value := range commandArray{
		command += value+" "
	}
	return
}

func exposedCli() {

	for {
		var commands[3]string
		fmt.Print(">")
		fmt.Scanf("%v %v %v", &commands[0], &commands[1], &commands[2])
		result := interpreter(commandBuilder(commands))
		fmt.Println(result)
		if result=="Bye" {
			break
		}
	}
}

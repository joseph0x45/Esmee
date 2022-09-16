package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func exposedCli() {
	var root string = ">"
	for {
		fmt.Print(root)
		reader := bufio.NewReader(os.Stdin)
		cmdReader, _ := reader.ReadString('\n')
		result := interpreter(strings.TrimSpace(cmdReader))
		if result == "Bye" {
			break
		}
		switch {
			case strings.Contains(result, "climb"):
				action := strings.Split(result, " ")[0]
				treeName := strings.Split(result, " ")[1]
				switch {
				case action=="climb":
					root += fmt.Sprintf("%v#", treeName)
					fmt.Printf("You are now in tree %v\n", treeName)
				case action=="climbdown":
					root = strings.Replace(root, fmt.Sprintf("%v#", treeName), "", 1)
				}
			case strings.Contains(result, "Error"):
				errorCode := strings.Split(result, " ")[1]
				fmt.Println(errorHandler(errorCode))
			default :
			fmt.Println(result)
		}
	}
}

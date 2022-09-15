package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func exposedCli() {

	for {
		// var commands[3]string
		fmt.Print(">")
		// fmt.Scanf("%v %v %v", &commands[0], &commands[1], &commands[2])
		reader := bufio.NewReader(os.Stdin)
		cmdReader, _ := reader.ReadString('\n')
		result := interpreter(strings.TrimSpace(cmdReader))
		fmt.Println(result)
		if result=="Bye" {
			break
		}
	}
}

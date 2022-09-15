package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func exposedCli() {

	for {
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		cmdReader, _ := reader.ReadString('\n')
		result := interpreter(strings.TrimSpace(cmdReader))
		fmt.Println(result)
		if result=="Bye" {
			break
		}
	}
}

package main

import (
	"fmt"
)

func exposedCli() {

	for {
		var cmd string
		fmt.Print(">")
		fmt.Scan(&cmd)
		result := interpreter(cmd)
		fmt.Println(result)
		if result=="Bye" {
			break
		}
	}
}

package main

import (
	
)

func interpreter(command string) (result string) {
	switch command {
	case "Help":
		result = "Welcome to EsmeeDB exposed CLI"
	case "help":
		result = "Welcome to EsmeeDB exposed CLI"
	case "exit":
		result = "Bye"
	}

	return
	
}
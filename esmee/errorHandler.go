package main

import (
	"fmt"
)

func errorHandler(code string) (meaning string) {
	switch code {
	case "TIT":
		meaning = "Can't create tree in a tree"
	case "IS":
		meaning = "Invalid syntax"
	case "UC":
		meaning = "Unrecognized command"
	case "ET":
		meaning = fmt.Sprintf("Can't create existing tree")
	}
	return
}
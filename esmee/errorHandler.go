package main

import (
)

func errorHandler(code string) (meaning string) {
	switch code {
	case "TIT":
		meaning = "Can't create tree in a tree"
	case "BOOT":
		meaning = "Can't create branch while not in a tree"
	case "IS":
		meaning = "Invalid syntax"
	case "UC":
		meaning = "Unrecognized command"
	case "ET":
		meaning = "Can't create existing tree"
	case "EB":
		meaning = "Can't create existing branch"
	case "EL":
		meaning = "Can't create existing leaf"
	case "ITOB":
		meaning = "Can't climb inexsting element"
	case "CDRK":
		meaning = "`down` is a reserved word. You can't create a tree named down"
	}
	return
}
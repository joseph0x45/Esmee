package main

func errorHandler(code string) (meaning string) {
	switch code {
	case "TIT":
		meaning = "Can't create tree in a tree"
	case "IS":
		meaning = "Invalid syntax"
	case "UC":
		meaning = "Unrecognized command"
	}
	return
}
package main

import (
	"fmt"
	"io/fs"
	"os"
)

const perm fs.FileMode = 0777

func createTree(treeName string) (status bool, message string) {
	err := os.Mkdir(treeName, perm)
	if err!= nil {
		status = false
		message = "Error ET"
		return
	}
	status = true
	message = fmt.Sprintf("Tree %v created", treeName)
	return
}

func createBranch(treeName string, branchName string) (status bool, message string) {
	path := treeName + "/" + branchName
	err := os.Mkdir(path, perm)
	if err!= nil{
		status = false
		message = "Error EB"
		return
	}
	status = true
	message = fmt.Sprintf("Branch %v created on tree %v", branchName, treeName)
	return
}

func removeTree(treeName string) (status bool, message string) {
	err := os.RemoveAll(treeName)
	if err!=nil{
		status = false
		message = "Can't remove tree"
		return
	}
	status = true
	message = fmt.Sprintf("Tree %v removed", treeName)
	return
}

func removeBranch(treeName string, branchName string) (status bool, message string)  {
	path := treeName + "/" + branchName
	err := os.RemoveAll(path)
	if err!=nil{
		status = false
		message = "Can't remove tree"
		return
	}
	status = true
	message = fmt.Sprintf("Tree %v removed", treeName)
	return
}

func createLeaf(treeName string, branchName string, leafName string) (status bool, message string) {
	path := treeName + "/" + branchName + "/" + leafName
	fmt.Printf("About to create file %v\n", path)
	err := os.WriteFile(path, []byte("Bruh"), 0777)
	if err!= nil{
		status = false
		message = "Can't create leaf"
		return
	}
	status = true
	message = fmt.Sprintf("Leaf %v created on tree %v", leafName, treeName)
	return
}
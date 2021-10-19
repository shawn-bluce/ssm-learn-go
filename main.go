package main

import (
	"bufio"
	"fmt"
	"os"
	"ssm-learn-go/structs"
	"ssm-learn-go/utils"
	"strings"
)


var linkedList structs.Node


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Student Score Manager")
	fmt.Println("---------------------------------------")

	linkedList = structs.Node{}

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "insert":
			utils.InsertStudent(&linkedList)
		case "delete":
			utils.DeleteStudent()
		case "modify":
			utils.ModifyStudent()
		case "search":
			utils.SearchStudent()
		case "list":
			utils.ListStudent(&linkedList)
		default:
			fmt.Println("invalid command, try again")
		}
	}
}

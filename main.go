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
		case "init":
			for i := 0; i < 20; i++ {
				utils.InsertStudent(&linkedList)
			}
			utils.ListStudent(&linkedList)
		case "insert":
			utils.InsertStudent(&linkedList)
		case "delete":
			utils.DeleteStudent()
		case "modify":
			utils.ModifyStudent()
		case "search":
			utils.SearchStudent(&linkedList)
		case "list":
			utils.ListStudent(&linkedList)
		case "quit":
			os.Exit(0)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("invalid command, try again")
		}
	}
}

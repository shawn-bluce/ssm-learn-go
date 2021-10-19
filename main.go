package main

import (
	"bufio"
	"fmt"
	"os"
	"ssm-learn-go/utils"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Student Score Manager")
	fmt.Println("---------------------------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text {
		case "insert":
			utils.InsertStudent()
		case "delete":
			utils.DeleteStudent()
		case "modify":
			utils.ModifyStudent()
		case "search":
			utils.SearchStudent()
		case "list":
			utils.ListStudent()
		default:
			fmt.Println("invalid command, try again")
		}
	}
}

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
		case "init":
			for i := 0; i < 20; i++ {
				utils.InsertStudent()
			}
			utils.ListStudent()
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
		case "quit":
			err := utils.GetDbConnection().Close()
			if err != nil {
				return
			}
			os.Exit(0)
		case "exit":
			err := utils.GetDbConnection().Close()
			if err != nil {
				return
			}
			os.Exit(0)
		default:
			fmt.Println("invalid command, try again")
		}
	}
}

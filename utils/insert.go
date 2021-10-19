package utils

import (
	"fmt"
	"ssm-learn-go/structs"
)

func InsertStudent(linkedList *structs.Node) {
	student := structs.Student{
		Name: "shawn",
		Gender: "M",
		OperationSystem: 100,
		Networking: 200,
		DataStruct: 300,
		Algorithms: 400,
		Golang: 500,
	}

	point := linkedList
	for true {
		if point.Student.Name == ""{
			point.Student = student
			break
		} else {
			if point.NextNode == nil {
				point.NextNode = &structs.Node{}
			} else {
				point = point.NextNode
			}
		}
	}

	fmt.Println("insert a student")
}

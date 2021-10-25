package utils

import (
	"fmt"
	"ssm-learn-go/structs"
)

func SearchStudent(linkedList *structs.Node) {
	fmt.Println("input name: ")
	point := linkedList

	var name string
	fmt.Scanln(&name)

	for {
		if point == nil {
			fmt.Println("student not found")
			return
		}

		if point.Student.Name == name {
			PrintStudentInfoByLine(point.Student)
			return
		} else {
			point = point.NextNode
		}
	}
}

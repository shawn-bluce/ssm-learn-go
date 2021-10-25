package utils

import (
	"fmt"
	"ssm-learn-go/structs"
)

func DeleteStudent(linkedList *structs.Node) {
	fmt.Println("input name: ")
	point := linkedList

	var name string
	fmt.Scanln(&name)

	for {
		if point == nil || point.NextNode == nil {
			fmt.Println("student not found")
			return
		} else {
			if point.NextNode.Student.Name == name {
				point.NextNode = point.NextNode.NextNode
				fmt.Println("delete done.")
				return
			} else {
				point = point.NextNode
			}
		}
	}
}

package utils

import (
	"fmt"
	"ssm-learn-go/structs"
)

func ListStudent(linkedList *structs.Node) {
	point := linkedList
	for true {
		if point == nil {
			break
		} else {
			fmt.Println(point.Student)
		}
		point = point.NextNode
	}
}

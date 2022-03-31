package utils

import (
	"fmt"
	"ssm-learn-go/structs"
)

func ListStudent() {
	db := GetDbConnection()
	query, err := db.Query("SELECT * FROM student")
	if err != nil {
		return
	}

	for query.Next() {
		var student structs.Student
		err = query.Scan(&student.Id, &student.Name, &student.OperationSystem, &student.Networking, &student.DataStruct, &student.Golang)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(fmt.Sprintf("%3d, %s, %2.2f, %2.2f, %2.2f, %2.2f", student.Id, student.Name, student.OperationSystem, student.Networking, student.DataStruct, student.Golang))
	}
}

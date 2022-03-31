package utils

import (
	"ssm-learn-go/structs"
)

func GetStudent(studentId string) structs.Student {
	db := GetDbConnection()

	var student structs.Student
	err := db.QueryRow("SELECT * FROM student WHERE id=?", studentId).Scan(
		&student.Id, &student.Name, &student.OperationSystem,
		&student.Networking, &student.DataStruct, &student.Golang)
	if err != nil {
		return structs.Student{}
	}

	return student
}

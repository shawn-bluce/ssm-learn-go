package utils

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"ssm-learn-go/structs"
)

func BuildStudentFromInput() structs.Student {
	return structs.Student{
		Name:            faker.FirstName(),
		OperationSystem: rand.Float32() * 100,
		Networking:      rand.Float32() * 100,
		DataStruct:      rand.Float32() * 100,
		Golang:          rand.Float32() * 100,
	}
}

func InsertStudent() {
	db := GetDbConnection()
	student := BuildStudentFromInput()
	values := fmt.Sprintf("'%s', %f, %f, %f, %f", student.Name, student.OperationSystem, student.Networking, student.DataStruct, student.Golang)
	db.Query(fmt.Sprintf("INSERT INTO student(name, operating_system, networking, data_struct, golang) VALUES (%s)", values))
}

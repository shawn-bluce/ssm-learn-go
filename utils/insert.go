package utils

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"math/rand"
	"ssm-learn-go/structs"
)

func BuildStudentFromInput(fakeMode bool) structs.Student{
	var name string
	var gender string
	var operationSystem float32
	var networking float32
	var dataStruct float32
	var algorithms float32
	var golang float32

	// only use fake mode, for testing and development
	if fakeMode {
		genderList := []string{"M", "F"}
		name = faker.FirstName()
		gender = genderList[rand.Intn(2)]
		operationSystem = rand.Float32() * 100
		networking = rand.Float32() * 100
		dataStruct = rand.Float32() * 100
		algorithms = rand.Float32() * 100
		golang = rand.Float32() * 100
	} else {
		fmt.Scanln(&name)
		fmt.Scanln(&gender)
		fmt.Scanln(&operationSystem)
		fmt.Scanln(&networking)
		fmt.Scanln(&dataStruct)
		fmt.Scanln(&algorithms)
		fmt.Scanln(&golang)
	}

	return structs.Student{
		Name:            name,
		Gender:          gender,
		OperationSystem: operationSystem,
		Networking:      networking,
		DataStruct:      dataStruct,
		Algorithms:      algorithms,
		Golang:          golang,
	}

}


func InsertStudent(linkedList *structs.Node) {
	student := BuildStudentFromInput(true)

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

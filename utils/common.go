package utils

import (
	"fmt"
	"ssm-learn-go/structs"
)

func PrintStudentInfoByLine(student structs.Student) {
	fmt.Println("\n\n\n\n\n============================")
	fmt.Printf("name            : %-6s \ngender          : %-2s \nOperationSystem : %4.2f\n", student.Name, student.Gender, student.OperationSystem)
	fmt.Printf("Networking      : %4.2f\nDataStruct      : %4.2f\nAlgorithms      : %4.2f\nGolang          : %4.2f\n", student.Networking, student.DataStruct, student.Algorithms, student.Golang)
	fmt.Println("============================")
}

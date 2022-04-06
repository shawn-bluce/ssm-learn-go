package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"ssm-learn-go/structs"
	"strings"
)

func HandleStudentList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	queryParm := strings.Split(httpRequest.URL.RawQuery, "&")
	whereContent := ""
	for _, parm := range queryParm {
		keyValue := strings.Split(parm, "=")
		whereContent = whereContent + " AND " + keyValue[0] + "=" + keyValue[1]
	}
	whereContent = strings.TrimLeft(whereContent, " AND ")

	db := GetDbConnection()
	sqlString := "SELECT * FROM student WHERE " + whereContent
	fmt.Println(sqlString)
	query, _ := db.Query(sqlString)

	for query.Next() {
		student := structs.Student{}

		query.Scan(&student.Id, &student.Name, &student.OperationSystem,
			&student.Networking, &student.DataStruct, &student.Golang)
		fmt.Println(student.Id, student.Name, student.OperationSystem,
			student.Networking, student.DataStruct, student.Golang)
	}
}

func HandleStudentSingle(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	studentId := regexp.MustCompile("[0-9]+").FindAllString(httpRequest.URL.Path, -1)[0]
	student := GetStudent(studentId)
	res, _ := json.Marshal(student)
	if student.Id > 0 {
		httpResponse.WriteHeader(200)
	} else {
		httpResponse.WriteHeader(404)
	}
	httpResponse.Header().Set("content-type", "application/json")
	httpResponse.Write(res)
}

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

func DeleteStudent(studentId string) bool {
	db := GetDbConnection()

	err, _ := db.Query("DELETE FROM student WHERE id=?", studentId)

	return err == nil
}

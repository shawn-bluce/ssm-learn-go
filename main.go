package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"ssm-learn-go/utils"
)

func handleStudentList(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(httpResponse, httpRequest.URL.RawQuery)
}

func handleStudentSingle(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	studentId := regexp.MustCompile("[0-9]+").FindAllString(httpRequest.URL.Path, -1)[0]
	student := utils.GetStudent(studentId)
	res, _ := json.Marshal(student)
	if student.Id > 0 {
		httpResponse.WriteHeader(200)
	} else {
		httpResponse.WriteHeader(404)
	}
	httpResponse.Header().Set("content-type", "application/json")
	httpResponse.Write(res)
}

func router(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	studentList := regexp.MustCompile(`^/student/\??`)
	studentSingle := regexp.MustCompile(`^/student/\d+/$`)

	switch {
	case studentSingle.MatchString(httpRequest.URL.Path):
		handleStudentSingle(httpResponse, httpRequest)
	case studentList.MatchString(httpRequest.URL.Path):
		handleStudentList(httpResponse, httpRequest)
	default:
		fmt.Println("do not case")
	}
}

func main() {
	http.HandleFunc("/", router)

	err := http.ListenAndServe(":5050", nil)
	if err != nil {
		return
	}
}

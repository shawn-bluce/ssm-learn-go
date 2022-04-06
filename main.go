package main

import (
	"fmt"
	"net/http"
	"regexp"
	"ssm-learn-go/utils"
)

func router(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	studentList := regexp.MustCompile(`^/student/?\??`)
	studentSingle := regexp.MustCompile(`^/student/\d+/$`)

	switch {
	case studentSingle.MatchString(httpRequest.URL.Path):
		utils.HandleStudentSingle(httpResponse, httpRequest)
	case studentList.MatchString(httpRequest.URL.Path):
		utils.HandleStudentList(httpResponse, httpRequest)
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

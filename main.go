package main

import (
	"fmt"
	"ssm-learn-go/structs"
)

func main() {
	me := structs.Student{Name: "Shawn", Gender: "M"}
	fmt.Println(me.Name, me.Gender)
}

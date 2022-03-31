package utils

import (
	"bufio"
	"fmt"
	"os"
)

func DeleteStudent() {
	fmt.Print("id -> ")
	id, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	db := GetDbConnection()
	sqlString := fmt.Sprintf("DELETE FROM student WHERE id=%s", id)
	db.Query(sqlString)
}

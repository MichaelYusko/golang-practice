package main

import (
	"fmt"
	"golang-practice/models"
)

func main() {
	user := models.NewUser("Mike", 150.00, "Tokyo")
	user.RaiseSalary(12.00)
	fmt.Println(user)
}

package main

import (
	"golang-practice/db"
	"fmt"
)


func main() {
	db := db.CreateDatabase("Development")
	fmt.Println(db.ShowName())
}

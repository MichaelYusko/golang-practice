package main

import (
	"fmt"
	"golang-practice/db"
)

func main() {
	db := db.CreateDatabase("Development")
	fmt.Println(db.ShowName())
}

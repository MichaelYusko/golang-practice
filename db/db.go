package db

import "fmt"

type Database struct {
	databaseName string
}

// Create a databae
func CreateDatabase(databaseName string) *Database {
	return &Database{databaseName}
}

// Show name of a database
func (d *Database) ShowName() string {
	return d.databaseName
}

// Create a table
// Not completed yet
func (d Database) CreateTable(name string) {
	fmt.Printf("%v was created.", name)
}

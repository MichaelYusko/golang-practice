package models

import (
	"math/rand"
)

type User struct {
	ID int
	Name string
	Salary float64
	City string
}

// Create a new user
func NewUser(name string, salary float64, city string) *User {
	return &User{
		ID: rand.Int(),
		Name: name,
		Salary: salary,
		City: city,
	}
}

// Raise a salary for a user
func (u *User) RaiseSalary(salary float64) {
	u.Salary += salary
}

package main

import (
	"fmt"
)

type User struct {
	Name string `json:"name" db:"name"`
}

var user User

type quiz string

func (g quiz) Render() {
	fmt.Println("Ola %v", user.Name)
}

var Quiz quiz

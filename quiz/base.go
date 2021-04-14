package quiz

import (
	"math/rand"
)

type Test struct {
	Title       string
	Cover       string
	Slug        string
	Description string
	Message     string
	Class       string
}

type User struct {
	Id    string
	Name  string
	Email string
	Photo string
}

func Choose(images []string) string {
	rnd := rand.Intn(len(images))
	return images[rnd]
}

package main

import (
	"fmt"
	"supertests/quiz"
)

var user quiz.User

func Render() {
	fmt.Printf("Ola %v", user.Name)
}

func Info() quiz.Test {
	return quiz.Test{
		Title:       "Como seria seus filhos Gêmeos?",
		Cover:       "como-seria-seus-filhos-gemeos/gemeoscapa.jpg",
		Slug:        "como-seria-seus-filhos-gemeos",
		Description: "Faça o teste e veja como seria seus filhos gêmeos",
		Message:     "[[nome]], olha como seria seus filhos, gostou?",
	}
}

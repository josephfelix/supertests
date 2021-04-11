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
		Title:       "Qual deveria ser seu salario?",
		Cover:       "qual-deveria-ser-seu-salario/capa.jpg",
		Slug:        "qual-deveria-ser-seu-salario",
		Description: "Faça o teste e descubra quanto deveria ser realmente seu salário",
		Message:     "[[nome]], é isto que você deveria estar ganhando. Gostou?",
	}
}

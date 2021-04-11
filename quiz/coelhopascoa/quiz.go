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
		Title:       "O que o coelhinho da Páscoa vai trazer pra você?",
		Cover:       "o-que-o-coelhinho-da-pascoa-vai-trazer-para-voce/coelinho230317.jpg",
		Slug:        "o-que-o-coelhinho-da-pascoa-vai-trazer-para-voce",
		Description: "Faça o teste e veja o que o coelhinho da Páscoa vai trazer para você!",
		Message:     "[[nome]], olha o que o coelho trouxe para você, gostou?",
	}
}

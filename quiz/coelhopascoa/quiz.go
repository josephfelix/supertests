package main

import (
	"fmt"
	"supertests/quiz"
)

var User quiz.User

func Render() (string, error) {
	return fmt.Sprintf("Ola %v", User.Name), nil
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

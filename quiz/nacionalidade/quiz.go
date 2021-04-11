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
		Title:       "Qual sua verdadeira nacionalidade?",
		Cover:       "qual-sua-verdadeira-nacionalidade/nacionalidadecapa.jpg",
		Slug:        "qual-sua-verdadeira-nacionalidade",
		Description: "Faça o teste e veja qual é a sua nacionalidade",
		Message:     "[[nome]], Essa é a sua verdadeira nacionalidade. Gostou?",
	}
}

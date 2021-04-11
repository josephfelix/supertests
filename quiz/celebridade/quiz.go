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
		Title:       "Qual celebridade se parece com você?",
		Cover:       "qual-celebridade-se-parece-com-voce/celebridadecapa.jpg",
		Slug:        "qual-celebridade-se-parece-com-voce",
		Description: "Descubra agora qual celebridade se parece mais com você!",
		Message:     "[[nome]], esta é a celebridade que você se parece. Gostou?",
	}
}

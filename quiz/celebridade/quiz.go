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
		Title:       "Qual celebridade se parece com você?",
		Cover:       "qual-celebridade-se-parece-com-voce/celebridadecapa.jpg",
		Slug:        "qual-celebridade-se-parece-com-voce",
		Description: "Descubra agora qual celebridade se parece mais com você!",
		Message:     "[[nome]], esta é a celebridade que você se parece. Gostou?",
	}
}

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
		Title:       "Quais Famosos estão te xavecando no whatsapp?",
		Cover:       "o-que-os-famosos-estao-falando-sobre-voce/whatsappcapa.jpg",
		Slug:        "quais-famosos-estao-te-xavecando-no-whatsapp",
		Description: "Faça o teste e veja quais famosos estão te enviando mensagens no WhatsAPP",
		Message:     "[[nome]], olha o que os famosos estão falando sobre você, gostou?",
	}
}

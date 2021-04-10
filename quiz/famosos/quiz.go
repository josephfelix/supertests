package main

import (
	"fmt"
	"supertests/models"
)

type Quiz string

var user models.User
var QuizInfo Quiz

func (g Quiz) Render() {
	fmt.Printf("Ola %v", user.Name)
}

func (g Quiz) Info() models.Test {
	return models.Test{
		Title:       "Quais Famosos estão te xavecando no whatsapp?",
		Cover:       "o-que-os-famosos-estao-falando-sobre-voce/whatsappcapa.jpg",
		Slug:        "quais-famosos-estao-te-xavecando-no-whatsapp",
		Description: "Faça o teste e veja quais famosos estão te enviando mensagens no WhatsAPP",
		Message:     "[[nome]], olha o que os famosos estão falando sobre você, gostou?",
	}
}

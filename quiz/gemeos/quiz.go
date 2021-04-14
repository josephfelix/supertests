package main

import (
	"fmt"
	"supertests/lib"
	"supertests/quiz"
)

var User quiz.User

var builder = lib.ImageBuilder{}

func Render() (string, error) {
	images := []string{
		"upload/como-seria-seus-filhos-gemeos/bebe01.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe02.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe03.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe04.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe05.jpg",
	}

	img := quiz.Choose(images)
	//picture := builder.Make(User.Photo)

	return fmt.Sprintf("Imagem: %v", img), nil
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

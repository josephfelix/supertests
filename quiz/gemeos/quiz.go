package main

import (
	"image"
	"supertests/lib"
	"supertests/quiz"
)

var User quiz.User

var builder = lib.ImageBuilder{}

func Render() (image.Image, error) {
	images := []string{
		"upload/como-seria-seus-filhos-gemeos/bebe01.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe02.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe03.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe04.jpg",
		"upload/como-seria-seus-filhos-gemeos/bebe05.jpg",
	}

	img, err := builder.Make(quiz.Choose(images))

	if err != nil {
		return nil, err
	}

	facebook, err := builder.Make(User.Photo)

	if err != nil {
		return nil, err
	}

	photoResized := builder.Resize(facebook, 260, 300)

	return builder.Insert(img, photoResized, 20, 80), nil
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

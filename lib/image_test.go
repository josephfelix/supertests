package lib

import (
	"supertests/quiz"
	"testing"
)

var builder = ImageBuilder{}

func Test_Image(t *testing.T) {
	images := []string{
		"../upload/como-seria-seus-filhos-gemeos/bebe01.jpg",
		"../upload/como-seria-seus-filhos-gemeos/bebe02.jpg",
		"../upload/como-seria-seus-filhos-gemeos/bebe03.jpg",
		"../upload/como-seria-seus-filhos-gemeos/bebe04.jpg",
		"../upload/como-seria-seus-filhos-gemeos/bebe05.jpg",
	}

	filename := quiz.Choose(images)

	if _, err := builder.ResolveImageFile(filename); err != nil {
		t.Fatalf(`Error: %v`, err)
	}

	if _, err := builder.Make(filename); err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}

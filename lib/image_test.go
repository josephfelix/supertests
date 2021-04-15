package lib

import (
	"supertests/quiz"
	"testing"
)

var builder = ImageBuilder{}
var images = []string{
	"../upload/como-seria-seus-filhos-gemeos/bebe01.jpg",
	"../upload/como-seria-seus-filhos-gemeos/bebe02.jpg",
	"../upload/como-seria-seus-filhos-gemeos/bebe03.jpg",
	"../upload/como-seria-seus-filhos-gemeos/bebe04.jpg",
	"../upload/como-seria-seus-filhos-gemeos/bebe05.jpg",
}

func Test_Image_ResolveImageFile(t *testing.T) {
	filename := quiz.Choose(images)

	if _, err := builder.ResolveImageFile(filename); err != nil {
		t.Fatalf(`Error: %v`, err)
	}

}

func Test_Make(t *testing.T) {
	filename := quiz.Choose(images)

	if _, err := builder.Make(filename); err != nil {
		t.Fatalf(`Error: %v`, err)
	}
}

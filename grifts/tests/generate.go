package tests

import (
	"fmt"
	"plugin"
	"supertests/lib"
	"supertests/models"
	"supertests/quiz"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/markbates/grift/grift"
)

func Generate(c *grift.Context) error {
	tests, err := lib.GetAllTestFiles()

	if err != nil {
		return err
	}

	for test := range tests {
		_, output := lib.AssembleQuizPath(test)

		p, err := plugin.Open(output)

		if err != nil {
			panic(err)
		}

		info, err := p.Lookup("Info")

		if err != nil {
			panic(err)
		}

		quizinfo := info.(func() quiz.Test)()

		models.DB.Transaction(func(tx *pop.Connection) error {
			quiztest := models.Test{
				Title:       quizinfo.Title,
				Slug:        quizinfo.Slug,
				Cover:       quizinfo.Cover,
				Description: quizinfo.Description,
				Message:     quizinfo.Message,
				Plugin:      output,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if err := tx.Create(&quiztest); err != nil {
				return err
			}

			fmt.Printf("Registrado quiz: %v\n", quizinfo.Title)

			return nil
		})
	}

	return nil
}

package tests

import (
	"fmt"
	"plugin"
	"supertests/lib"
	"supertests/models"

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

		fmt.Println(info.(func() models.Test)())
	}

	return nil
}

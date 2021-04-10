package grifts

import (
	"supertests/grifts/tests"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("tests", func() {
	grift.Desc("generate", "Generate tests")
	grift.Add("generate", tests.Generate)

	grift.Desc("compile", "Compile tests")
	grift.Add("compile", tests.Compile)

	grift.Desc("all", "Generate and compile tests")
	grift.Add("all", func(c *grift.Context) error {
		err := grift.Run("tests:generate", c)
		if err != nil {
			return err
		}

		err = grift.Run("tests:compile", c)
		if err != nil {
			return err
		}

		return nil
	})
})

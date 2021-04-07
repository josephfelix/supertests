package grifts

import (
	"fmt"

	"github.com/markbates/grift/grift"
)

var _ = grift.Namespace("tests", func() {
	grift.Desc("generate", "Generate tests")
	grift.Add("generate", func(c *grift.Context) error {
		fmt.Println("Teste")
		return nil
	})
})

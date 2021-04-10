package tests

import (
	"fmt"
	"log"
	"supertests/lib"

	"github.com/markbates/grift/grift"
)

func Compile(c *grift.Context) error {
	tests, err := lib.GetAllTestFiles()

	if err != nil {
		return err
	}

	for test := range tests {
		entry, output := lib.AssembleQuizPath(test)

		command := fmt.Sprintf("go build -buildmode=plugin -o %v %v", output, entry)

		_, stderr, err := lib.Shellout(command)

		if stderr != "" {
			log.Fatalln(stderr)
			return nil
		}

		if err != nil {
			return err
		}

		fmt.Printf("Compiled test: %v\n", entry)
	}

	return nil
}

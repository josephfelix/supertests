package lib

import (
	"os"
	"path/filepath"
	"plugin"
)

func GetAllTestFiles() (<-chan string, error) {
	basedir := "quiz"

	f, err := os.Open(basedir)

	if err != nil {
		return nil, err
	}

	files, err := f.Readdir(-1)
	f.Close()

	if err != nil {
		return nil, err
	}

	tests := make(chan string)

	go func() {
		for _, file := range files {
			filename, _ := filepath.Abs(basedir + "/" + file.Name() + "/quiz.go")

			if _, err := os.Stat(filename); err == nil {
				tests <- file.Name()
			}
		}

		close(tests)
	}()

	return tests, nil
}

func AssembleQuizPath(test string) (string, string) {
	compiledir, _ := filepath.Abs("plugins")

	outputfile := compiledir + "/" + test + ".so"
	entryfile, _ := filepath.Abs("quiz/" + test + "/quiz.go")

	return entryfile, outputfile
}

func QuizLoad(filename string) *plugin.Plugin {
	quiz, err := plugin.Open(filename)

	if err != nil {
		panic(err)
	}

	return quiz
}

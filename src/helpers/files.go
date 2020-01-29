package helpers

import (
	"io/ioutil"
	"log"
)

type files struct{}

var (
	Files files
)

func (f *files) ReadFile(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	return text
}

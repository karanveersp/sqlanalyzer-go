package util

import (
	"io/ioutil"
)

// ReadFileAsString returns all contents of the file as a single string.
func ReadFileAsString(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(content)
}

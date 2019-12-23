package file

import (
	"os"
	"io/ioutil"
)

func GetFileContent(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return b
}

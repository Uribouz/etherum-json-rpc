package mockdata

import (
	"fmt"
	"os"
)

const FILE_NAME = "mockdata/addresses.json"

type fileReader struct {
	filename string
}

func NewFileReader() fileReader {
	return fileReader{filename: FILE_NAME}
}

func (f fileReader) GetData() (string, error) {
	data, err := os.ReadFile(f.filename)
	if err != nil {
		return "",fmt.Errorf("cannot ReadFile: %v", err)
	}
	return string(data), nil
}
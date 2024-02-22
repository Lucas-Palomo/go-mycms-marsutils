package internal

import (
	"io"
	"os"
)

var (
	Open     = os.Open
	ReadAll  = io.ReadAll
	OpenFile = os.OpenFile
)

func ReadFile(filepath string) ([]byte, error) {
	file, err := Open(filepath)

	if err != nil {
		return nil, err
	}

	bytes, err := ReadAll(file)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	return bytes, nil
}

func WriteFile(filepath string, content []byte) error {
	file, err := OpenFile(filepath, os.O_SYNC|os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0666)

	if err != nil {
		return err
	}

	_, err = file.Write(content)

	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}

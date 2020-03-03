package utils

import (
	"bufio"
	"fmt"
	"os"
)

func CheckFileExists(filename string) bool{
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
			return false
		}
	}
	return true
}


func ReadFileContent(filename  string) ([]string, error) {
	var contentList []string
	fileptr, err := os.Open(filename)

	if err != nil {
		return nil, err
	}

	defer func() {
		err = fileptr.Close()
	}()

	fileContent:= bufio.NewScanner(fileptr)
	for fileContent.Scan() {
		contentList = append(contentList, fileContent.Text())
	}

	return contentList, nil
}

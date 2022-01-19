package goreloaded

import (
	"fmt"
	"os"
)

func ReadFile(file_input string) string {

	file, err := os.Open(file_input)
	if err != nil {
		fmt.Println(err)
	}

	aFile, _ := file.Stat()

	size := aFile.Size()

	arr := make([]byte, size)

	file.Read(arr)

	defer file.Close()

	return string(arr)

}

package goreloaded

import (
	"fmt"
	"os"
)

func ReadFile( filename string){
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(content))
}
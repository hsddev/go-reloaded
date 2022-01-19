package main

import (
	"fmt"
	"goreloaded"
	"os"
	"strings"
)

func main() {
	arguments := os.Args[1:]

	file_input := arguments[0]
	file := goreloaded.ReadFile(file_input)
	filesplit := strings.Split(file, " ")  //returns slice of strings
	filehexa := goreloaded.HexaConvert(filesplit) //basic (cap/bin/hex...) fixed
	filehexanumb := goreloaded.Hexnumbered(filehexa)
	filepunctone := goreloaded.Punctone(filehexanumb) // to delete
	filepuncttwo := goreloaded.Puncttwo(filepunctone)
	fileatoan := goreloaded.Atoan(filepuncttwo)

	man := os.WriteFile(arguments[1], []byte(fileatoan), 0644)
	if man != nil {
		panic(man)
	}
	// fmt.Printf("\nFINAL OUTPUT : %v", filehexa)
	// fmt.Printf("\nFINAL OUTPUT : %v", filehexanumb)
	// fmt.Printf("\nFINAL OUTPUT : %v", filepunctone)
	// fmt.Printf("\nFINAL OUTPUT : %v", filepuncttwo)
	fmt.Printf("\nFINAL OUTPUT : %v", fileatoan)

}

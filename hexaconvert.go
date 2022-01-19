package goreloaded

import (
	"fmt"
	"strconv"
	"strings"
)

func HexaConvert(n []string) []string {

	var empty []string

	for i := 0; i < len(n); i++ {
		if n[i] == "(hex)" { //if "(hex)" is found
			num, _ := strconv.ParseInt(n[i-1], 16, 64) //fix preceeding value
			empty[len(empty)-1] = fmt.Sprint(num)
			
		}
		

		if n[i] == "(bin)" {
			num, _ := strconv.ParseInt(n[i-1], 2, 64)
			empty[len(empty)-1] = fmt.Sprint(num)

		}

		if n[i] == "(up)" {
			empty[len(empty)-1] = strings.ToUpper(n[i-1])

		}

		if n[i] == "(low)" {
			empty[len(empty)-1] = strings.ToLower(n[i-1])

		}

		if n[i] == "(cap)" {
			empty[len(empty)-1] = strings.Title(empty[i-1])

		}

		if n[i] != "(hex)" && n[i] != "(bin)" && n[i] != "(up)" && n[i] != "(low)" && n[i] != "(cap)" {
			empty = append(empty, n[i])

		}
	}
	fmt.Printf("%#v", empty)
	return empty

}

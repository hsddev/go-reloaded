package goreloaded

import (
	"strings"
)

func Punctone(n []string) string {

	sjoined := strings.Join(n, " ")

	//fmt.Println("\nThis is sjoined:", sjoined)

	nrune := []rune(sjoined)

	//fmt.Println("\n This is nrune", nrune)

	for i := 0; i < len(nrune); i++ {

		if IsPunc(nrune[i]) && nrune[i-1] == 32 {
			nrune[i], nrune[i-1] = nrune[i-1], nrune[i]

		}
	}
	return StandardizeSpaces(string(nrune))
}

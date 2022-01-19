package goreloaded

import "strings"

func Atoan(s string) string {

	sos := strings.Split(s, " ")

	for i := 0; i < len(sos)-1; i++ {

		if sos[i] == "a" && IsVowelh(sos[i+1]) {
			sos[i] = "an"
		}

		if sos[i] == "A" && IsVowelh(sos[i+1]) {
			sos[i] = "An"

		}

	}
	return strings.Join(sos, " ")

}

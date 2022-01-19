package goreloaded

import (
	"strings"
)

func Hexnumbered(s []string) []string {

	pd := 0
	// var empty []string

	for i := 0; i < len(s); i++ {

		if s[i] == "(up," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				s[i-j] = strings.ToUpper(s[i-j])
			}

			s = append(s[:i], s[i+2:]...)
			// fmt.Printf("\nEMPTY AFTER UP: %v\n", s)

		}

		if s[i] == "(low," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				s[i-j] = strings.ToLower(s[i-j])
			}
			s = append(s[:i], s[i+2:]...)
			// fmt.Printf("\nEMPTY AFTER LOW: %v\n", s)
		}

		if s[i] == "(cap," {
			pd = TrimAtoi(s[i+1])

			for j := 1; j <= pd; j++ {
				s[i-j] = strings.Title(s[i-j])
			}

			s = append(s[:i], s[i+2:]...)
			// fmt.Printf("\nEMPTY AFTER CAP: %v\n", s)

		}

	}
	return s
}

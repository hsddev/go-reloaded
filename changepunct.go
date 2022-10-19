package reloaded

import "strings"

// fix punctuation position
func ChangePunctuationPosition(el string) string{
	result := []rune(el)
		
	for i:=0; i<len(result)-1; i++{
		if IsPunc(result[i+1]) && result[i] == ' '{
			result[i+1], result[i] = result[i], result[i+1]
		}
	}
	return strings.Join(strings.Fields(string(result)), " ")
		}
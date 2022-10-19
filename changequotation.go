package reloaded

import "strings"

func ChangeQuotationMark(el string) string{
	result := []rune(el)
	for i, k:= 0, len(result)-1; i<k && i!=k; i, k = i+1, k-1{
		if result[i] == '\'' && result[i+1] == ' '{
			result[i], result[i+1] = result[i+1], result[i]
		}
		if result[k] == '\'' && result[k-1] == ' '{
			result[k], result[k-1] = result[k-1], result[k]
		}
		}
	return strings.Join(strings.Fields(string(result)), " ")
}
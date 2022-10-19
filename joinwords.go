package reloaded

import "strings"

// Join the words in one sentence and add a space between each word
func JoinWords(el []string) string {
	return strings.Join(strings.Fields(strings.Join(el," ")), " ")
}
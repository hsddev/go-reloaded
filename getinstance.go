package reloaded

import "strings"

// Split the content into words and put them in a slice
func GetInstance(content string) []string {
	return strings.Split(content, " ")
}
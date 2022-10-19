package reloaded

// Check if vowel
func IsVowel(el string) bool {
	var vowels = []rune{'A','E','I','O','U','a','e','i','o','u','h'}
	for _,v:=range vowels{
		if (v == rune(el[0])){
			return true
		}
	}
	return false
}
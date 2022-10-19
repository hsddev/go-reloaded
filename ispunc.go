package reloaded

// Check if punctuation
func IsPunc(el rune) bool {
	switch (el){
		case '.',
		',',
		'!',
		':',
		';',
		'?': return true
	}
	return false
}
package goreloaded

func IsVowelh(n string) bool {

	var result bool

	for i := 0; i < len(n); i++ {
		if n[0] == 'A' || n[0] == 'E' || n[0] == 'I' || n[0] == 'O' || n[0] == 'U' || n[0] == 'a' || n[0] == 'e' ||
			n[0] == 'i' || n[0] == 'o' || n[0] == 'u' || n[0] == 'h' {
			result = true

		} else {
			result = false
		}
	}

	return result

}

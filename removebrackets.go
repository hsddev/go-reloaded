package goreloaded

func RemoveBrackets(n string) string {

	var opening int
	var closing int

	for index, char := range n {
		if char == '(' {
			opening = index
		} else if char == ')' {
			closing = index + 1

		}
	}

	toAdd := closing - opening

	var empty string

	for i := 0; i < len(n); i++ {
		if i == opening {
			i += toAdd
		} else {
			empty += string(n[i])
			//empty = append(empty, string(n[i]))
		}
	}
	return empty
}

package goreloaded

func Reverse(n []string) []string {

	reversed := []string{}

	for _, v := range n {
		reversed = append([]string{v}, reversed...)
		// result = string(v) + result
	}
	return reversed //fmt.Sprint(reserved)
}

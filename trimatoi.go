package goreloaded

func TrimAtoi(s string) int {
	neg := false       // intialise neg as false
	slice := []rune(s) // slice string into a rune to manipulate it
	trim := 0          // initialise trim as 0
	for i := 0; i < len(slice); i++ {

		if !neg && trim == 0 && slice[i] == '-' {
			neg = true
		}
		if slice[i] >= '0' && slice[i] <= '9' {
			trim *= 10
			trim += int(slice[i] - 48)
		}
	}
	if neg {
		return trim * -1
	}
	return trim
}

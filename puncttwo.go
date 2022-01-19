package goreloaded

func Puncttwo(n string) string {

	nrune := []rune(n)
	count := 0

	// 39 = '  32 = space
	//fmt.Println(string(nrune))

	for i := 0; i <= len(nrune)-1; i++ {

		if nrune[i] == 39 {
			count++
			//fmt.Println(count)

			if nrune[i] == 39 && count%2 != 0 && nrune[i+1] == 32 {
				nrune[i], nrune[i+1] = nrune[i+1], nrune[i]
				i++
			}

			if nrune[i] == 39 && nrune[i-1] == 32 && count%2 == 0 {
				nrune[i], nrune[i-1] = nrune[i-1], nrune[i]
				//fmt.Println("Test")
			}
		}
		//fmt.Printf("\nString Nrune : %v", string(nrune))

	}
	return StandardizeSpaces(string(nrune))
}

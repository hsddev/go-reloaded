package reloaded

// a to an
func AtoAn(el []string)[]string{
	for i:=0;i<len(el);i++{
		switch(el[i]){
		case "a":
		if(IsVowel(el[i+1])){el[i]="an"}
		case "A":
		if(IsVowel(el[i+1])){el[i]="an"}
		}
	}
	return el
}
package reloaded

import (
	"strconv"
	"strings"
)

// Advanced convert (cap, number)..
func AdvancedConvert(el []string)[]string{
	var id int
	for i:=0;i<len(el);i++{
		switch(el[i]){
		case "(up,":
			id,_ = strconv.Atoi(strings.Trim(el[i+1], ")"))
			for j := 1; j <= id; j++ {
				el[i-j] = strings.ToUpper(el[i-j])
			}
			el = append(el[:i], el[i+2:]...)
		case "(low,":
			id,_ = strconv.Atoi(strings.Trim(el[i+1], ")"))
			for j := 1; j <= id; j++ {
				el[i-j] = strings.ToLower(el[i-j])
			}
			el = append(el[:i], el[i+2:]...)
	    case "(cap,":
			id,_ = strconv.Atoi(strings.Trim(el[i+1], ")"))
			for j := 1; j <= id; j++ {
				el[i-j] = strings.Title(el[i-j])
			}
			el = append(el[:i], el[i+2:]...)
		}
	}
	return el
}
package reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// Basic convert (cap) - (hex) - (bin)..s
func BasicConvert(el []string) []string{
	var converted []string
	
	for i:=0;i<len(el);i++{
		switch (el[i]){
		case "(hex)":
			num, _ := strconv.ParseInt(el[i-1], 16, 64)
			converted[len(converted)-1] = fmt.Sprint(num)
		case "(bin)":
			num, _ := strconv.ParseInt(el[i-1], 2, 64)
			converted[len(converted)-1] = fmt.Sprint(num)
		case "(up)":
			converted[len(converted)-1] = strings.ToUpper(el[i-1])
		case "(low)":
			converted[len(converted)-1] = strings.ToLower(el[i-1])
	    case "(cap)":
			converted[len(converted)-1] = strings.Title(converted[i-1])
		default:
			converted = append(converted, el[i])	
		}
		}
	return converted
}
package goreloaded

import (
	"fmt"
	"strconv"
)

func BinaryConvert(n string) string {

	num, err := strconv.ParseInt(n, 2, 64)
	if err != nil {
		panic(err)
	}
	return fmt.Sprint(num)

}

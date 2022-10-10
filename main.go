package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// fix punctuation position
func changePunctuationPosition(el string) string{
result := []rune(el)
	
	// return el
	for i:=0;i<len(el);i++{
		if(i<=len(el)-2){
			if (el[i]==' ' && isPunc(rune(el[i+1]))){
				result[i], result[i+1]= result[i+1], result[i]
			}
			
		}else {
			if (el[i+1]==' '{
				result[i+1] = 'x'
			}
		}

		}
		return strings.TrimSpace(string(result))
	}
	


// Join the words in one sentence and add a space between each word
func joinWords(el []string) string {
	return strings.Join(strings.Fields(strings.Join(el," ")), " ")
}

// a to an
func atoAn(el []string)[]string{
	for i:=0;i<len(el);i++{
		switch(el[i]){
		case "a":
		if(isVowel(el[i+1])){el[i]="an"}
		case "A":
		if(isVowel(el[i+1])){el[i]="an"}
		}
	}
	return el
}

// Advanced convert (cap, number)..
func advancedConvert(el []string)[]string{
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

// Basic convert (cap) - (hex) - (bin)..s
func basicConvert(el []string) []string{
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

// Check if punctuation
func isPunc(el rune) bool {
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

// Check if vowel 
func isVowel(el string) bool {
	var vowels = []rune{'A','E','I','O','U','a','e','i','o','u','h'}
	for _,v:=range vowels{
		if (v == rune(el[0])){
			return true
		}
	}
	return false
}

// Split the content into words and put them in a slice
func getInstance(content string) []string {
	return strings.Split(content, " ")
}

func main(){
	// Get arguments
	// arguments := os.Args;

	// reading the file
	content, _ := ioutil.ReadFile("sample.txt")

	// slice := getInstance(string(content));
	//joinWords(advancedConvert(basicConvert
	slice1:= changePunctuationPosition(string(content))


	// writing into the file
   fmt.Println(slice1)
	
}
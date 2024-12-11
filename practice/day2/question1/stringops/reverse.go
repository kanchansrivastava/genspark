package stringops

import "strings"

func reverseString(s string) string{
	//trim
	return strings.Trim(s, " ") 
}
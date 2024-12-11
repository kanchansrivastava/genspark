package stringops


func ReverseAndUppercase(s1 string, s2 string) string{
	result := reverseString(s1) + toUpperCase(s2)
	return result
}


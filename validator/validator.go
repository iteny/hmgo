package validator

import "strings"

// IsEmail check if the string is an email.
func Email(str string) bool {
	// TODO uppercase letters are not supported
	return hmEmail.MatchString(str)
}

// func Int(str string, start int, end int) bool {
// 	if start != 0 && end != 0 {
//
// 	}
// }
func ByteLength(str string, min, max int) bool {
	return strings.Count(str, "")-1 >= min && strings.Count(str, "")-1 <= max
}

package validator

import (
	"strings"
	"time"
)

const RF3339Js = "2006-01-02 15:04:05"
const RF3339T = "2006-01-02T15:04:05"

// Email check if the string is an email.
// 检查字符串是否是电子邮件
func Email(str string) bool {
	// TODO uppercase letters are not supported
	return hmEmail.MatchString(str)
}

// Length check if the string's length (in unicode) falls in a range.
// Length检查字符串的长度（以unicode为单位）是否在一个范围内
func Length(str string, min, max int) bool {
	return strings.Count(str, "")-1 >= min && strings.Count(str, "")-1 <= max
}

// ByteLength check if the string's length (in bytes) falls in a range.
// ByteLength检查字符串的长度（以字节为单位）是否在一个范围内
func ByteLength(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// Null check if the string is null.
// Null检查字符串是否为空
func Null(str string) bool {
	return len(str) == 0
}

// Numeric check if the string contains only numbers. Empty string is invalid.
// Numeric检查字符串是否只包含数字。空字符串无效。
func Numeric(str string) bool {
	if Null(str) {
		return false
	}
	return hmNumeric.MatchString(str)
}

// English check if the string contains only letters (a-zA-Z). Empty string is invalid.
// English检查如果字符串只包含字母（a-za-z）,空字符串无效
func English(str string) bool {
	if Null(str) {
		return false
	}
	return hmEnglish.MatchString(str)
}

// EnglishNumeric check if the string contains only letters and numbers. Empty string is invalid.
// EnglishNumeric检查字符串是否只包含字母和数字。空字符串无效
func EnglishNumeric(str string) bool {
	if Null(str) {
		return false
	}
	return hmEnglishNumeric.MatchString(str)
}

// IsTime check if string is valid according to given format
func Time(str string, format string) bool {
	_, err := time.Parse(format, str)
	return err == nil
}

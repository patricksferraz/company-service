package utils

import (
	"bytes"
	"os"
	"regexp"
	"unicode"
)

func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func CleanNonDigits(str *string) {
	buf := bytes.NewBufferString("")
	for _, r := range *str {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*str = buf.String()
}

func IsClock(str *string) bool {
	match, _ := regexp.MatchString("^([0-1]?[0-9]|[2][0-3]):?([0-5][0-9])(:?[0-5][0-9])?$", *str)
	return match
}

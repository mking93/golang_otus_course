package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {

	result := ""
	strBefore := ""
	for i := 0; i < len(str); i++ {
		if (i == 0 || (i+1 < len(str) && isInt(string(str[i+1])))) && isInt(string(str[i])) {
			result = ""
			return "", ErrInvalidString
		}

		v, _ := strconv.Atoi(string(str[i]))
		if isInt(string(str[i])) && v > 0 {

			if string(str[i-1]) == "\n" {
				result = result[:len(result)-1]
				result += strings.Repeat("\\n", v)
			} else {
				result += strings.Repeat(strBefore, v-1)
			}
		}

		if isInt(string(str[i])) && v <= 0 {
			result = result[:len(result)-1]
		}

		if !isInt(string(str[i])) {
			strBefore = string(str[i])
			result += strBefore
		}
	}
	return result, nil
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

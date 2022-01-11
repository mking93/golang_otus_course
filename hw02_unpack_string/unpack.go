package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var err = errors.New("invalid string")

func Unpack(str string) (string, error) {

	result := ""
	strBefore := ""
	for i := 0; i < len(str); i++ {
		if (i == 0 || (i+1 < len(str) && isInt(string(str[i+1])))) && isInt(string(str[i])) {
			result = ""
			return "", err
		}

		if isInt(string(str[i])) {
			v, _ := strconv.Atoi(string(str[i]))
			if v > 0 {
				if string(str[i-1]) == "\n" {
					result = result[:len(result)-1]
					result += strings.Repeat("\\n", v)
				} else {
					result += strings.Repeat(strBefore, v-1)
				}
			} else {
				result = result[:len(result)-1]
			}
		} else {
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

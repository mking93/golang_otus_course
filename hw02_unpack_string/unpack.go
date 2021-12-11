package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(str string) (string, error) {
	var err = errors.New("")
	result := ""
	strBefore := ""
	for i := 0; i < len(str); i++ {
		if (i == 0 || (i+1 < len(str) && isInt(string(str[i+1])))) && isInt(string(str[i])) {
			result = ""
			err = errors.New("Invalid string")
			break
		} else if isInt(string(str[i])) {
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
	return result, err
}

func main() {
	fmt.Println(Unpack("a4bc2d5e"))
	fmt.Println(Unpack("abcd"))
	fmt.Println(Unpack("3abc"))
	fmt.Println(Unpack("45"))
	fmt.Println(Unpack("aaa10b"))
	fmt.Println(Unpack("aaa0b"))
	fmt.Println(Unpack(""))
	fmt.Println(Unpack("d\n5abc"))
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// Problem
// https://www.codewars.com/kata/52b757663a95b11b3d00062d

package main

import (
	"strings"
)

func toWeirdCase(str string) string {
	weirdCase := ""
	evenCounter := 0
	for i := 0; i < len(str); i++ {
		evenCounter++
		if string(str[i]) == " " {
			evenCounter--
		}
		if i%2 == 0 {
			weirdCase += strings.ToUpper(string(str[i]))
		} else {
			weirdCase += strings.ToLower(string(str[i]))
		}
	}
	return weirdCase
}

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
		if evenCounter%2 == 0 {
			weirdCase += strings.ToUpper(string(str[i]))
		} else {
			weirdCase += strings.ToLower(string(str[i]))
		}
		if string(str[i]) == " " {
			evenCounter = 0
			continue
		}
		evenCounter++
	}
	return weirdCase
}

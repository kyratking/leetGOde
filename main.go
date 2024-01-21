package main

import "fmt"

func main() {
	fmt.Println(RomanToInt("VII"))
	fmt.Println(RomanToInt("MMXX"))
	fmt.Println(RomanToInt("MMXXVI"))
}

/*
dotest(0, "0")
dotest(1, "01")
dotest(2, "010")
dotest(3, "01001")
dotest(5, "0100101001001")
*/

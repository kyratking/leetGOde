package main

import "fmt"

func main() {
	fmt.Println("a"[0])
	fmt.Println("z"[0])
	fmt.Println("A"[0])
	fmt.Println("Z"[0])
}

/*
dotest(0, "0")
dotest(1, "01")
dotest(2, "010")
dotest(3, "01001")
dotest(5, "0100101001001")
*/

package main

import "fmt"

func main() {
	fmt.Println(theMaximumAchievableX(4, 1))
	fmt.Println(theMaximumAchievableX(3, 2))
	fmt.Println(theMaximumAchievableX(5, 4))
	fmt.Println(theMaximumAchievableX(7, 3))
}

/*
dotest(0, "0")
dotest(1, "01")
dotest(2, "010")
dotest(3, "01001")
dotest(5, "0100101001001")
*/

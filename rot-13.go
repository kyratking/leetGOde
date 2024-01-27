package main

func Rot13(msg string) string {
	a := 97
	z := 122
	A := 65
	Z := 90
	result := ""
	for i := 0; i < len(msg); i++ {
		bytecode := msg[i]
		if bytecode >= 97 && bytecode <= 122 {
			newCode := bytecode + 13
			if bytecode > 122 {
				newCode = z - newCode
			}
		} else if bytecode >= 65 && bytecode <= 90 {

		} else {
			result += string(bytecode)
		}
	}
	return result
}

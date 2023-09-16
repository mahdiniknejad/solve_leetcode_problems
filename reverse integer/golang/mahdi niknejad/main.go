package main

import (
	"math"
	"strconv"
)

func main() {
	var number int = 54321
	reverse(number)
}

func reverse(x int) int {
	a := 1
	if x >= 0 {
		a = +1
	} else {
		a = -1
		x *= a
	}
	s := strconv.Itoa(x)
	s = reverseString(s)
	number, _ := strconv.Atoi(s)
	number *= a
	if number > int(math.Pow(2, 31)) || number < int(math.Pow(-2, 31)) {
		return 0
	} else {
		return number
	}

}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

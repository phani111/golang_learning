// This code demonstrates how to work with runes (Unicode code points) in Go.
// The `len()` function returns the number of bytes in a string, which may not be the same as the number of runes (characters).
// The first string "hello" has 5 bytes, one for each ASCII character.
// The second string "汉" has 3 bytes, as it represents a single Chinese character.
// The third string is created by explicitly specifying the byte values for the Chinese character "汉", and also has 3 bytes.
package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(len(s))

	s = "汉"
	fmt.Println(len(s))

	s = string([]byte{0xE6, 0xB1, 0x89})
	fmt.Println(len(s))
}

// getIthRune returns the rune (Unicode code point) at the specified index in the given string.
// If the index is out of bounds, it returns -1.
// The use of -1 as a return value for out-of-range indices is a common practice in Go, serving as an easily recognizable error condition.

func getIthRune(largeString string, i int) rune {
	for idx, v := range largeString {
		if idx == i {
			return v
		}
	}
	return -1
}

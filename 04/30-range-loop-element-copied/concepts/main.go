// The main function demonstrates the use of the range loop to iterate over a slice of strings.
//
// The first loop iterates over the slice and prints the index and value of each element.
// The second loop iterates over the slice and prints the value of each element.
package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	for i, v := range s {
		fmt.Printf("index=%d, values=%s\n", i, v)
	}

	for _, v := range s {
		fmt.Printf("value=%s\n", v)
	}
}

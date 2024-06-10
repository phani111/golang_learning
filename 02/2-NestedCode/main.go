/*
join concatenates the two input strings s1 and s2, and returns the result. If the

	concatenated string is longer than the specified max length, it will be truncated
	to the max length. If either s1 or s2 is empty, an error is returned.

func join(s1, s2 string, max int) (string, error)

	correctUsage is a variant of the join function that handles empty input strings
	more gracefully. If s1 or s2 is empty, it will return the non-empty string
	truncated to the max length, rather than returning an error.

func correctUsage(s1, s2 string, max int) (string, error)

	concatenate simply concatenates the two input strings s1 and s2, and returns the
	result. This is a helper function used by join and correctUsage.

func concatenate(s1, s2 string) (string, error)
*/
package main

import (
	"errors"
	"log"
)

func main() {
	r, err := join("Hello, ", "World!", 10)
	if err != nil {
		panic(err)
	}
	log.Println(r)

	r, err = correctUsage("Hello, ", "World!", 10)
	if err != nil {
		panic(err)
	}
	log.Println(r)
}

func join(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")

	} else {
		if s2 == "" {
			return "", errors.New("s2 is empty")
		} else {
			concat, err := concatenate(s1, s2)
			if err != nil {
				return "", err
			} else {
				if len(concat) > max {
					return concat[:max], nil
				} else {
					return concat, nil
				}
			}
		}
	}
}

func correctUsage(s1, s2 string, max int) (string, error) {
	if s1 == "" {
		return "", errors.New("s1 is empty")

	}

	if s2 == "" {
		return "", errors.New("s2 is empty")

	}
	concat, err := concatenate(s1, s2)
	if err != nil {
		return "", err
	}

	if len(concat) > max {
		return concat[:max], nil
	}
	return concat, nil
}

func concatenate(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

package main

import (
	"regexp"
	"strings"
)

func stripPhoneNumber(phoneNumber string) string {
	return strings.Map(func(r rune) rune {
		if r == ' ' || r == '-' {
			return -1
		} else {
			return r
		}
	}, phoneNumber)
}

func phoneNumber() {

	// Phone number regex
	pattern, err := regexp.Compile(`(0|\+33)[1-9]{9}`)
	if err != nil {
		panic(err)
	}

	// Test
	phoneNumbers := []string{
		"0123456789",
		"012345678",
		"123456789",
		"0123456789",
		"+33123456789",
		"+330123456789",
	}

	for _, phoneNumber := range phoneNumbers {
		if pattern.MatchString(stripPhoneNumber(phoneNumber)) {
			println(phoneNumber, "YAY")
		} else {
			println(phoneNumber, "NAY")
		}
	}
}

func main() {
	phoneNumber()
}

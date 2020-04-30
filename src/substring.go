package main2

import (
	"fmt"
	"strings"
)

func main2() {
	note := "helloGolang"
	fmt.Println("This is subString:", subString1(note))
}

func subString1(note string) string {

	var a, b string

	fmt.Println(len(note))

	fmt.Println("enter two strings to perform operations a, b:")
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	firstString := strings.Index(note, a)
	if firstString == -1 {
		return ""
	}
	endString := strings.Index(note, b)
	if endString == -1 {
		return ""
	}
	subString := firstString + len(a)
	if subString >= endString {
		return ""
	}
	return note[subString:endString]
}

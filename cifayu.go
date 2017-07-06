package main

import "fmt"

func main() {
	x := "hello"
	for _, v := range x {
		v := v + 'A' - 'a'
		fmt.Printf("%c", v) // "HELLO" (one letter per iteration)
	}
}

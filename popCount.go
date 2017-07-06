package main

import "fmt"

func main() {
	fmt.Println(PopCount(0xfe))
}

func PopCount(x uint64) int {
	count := 0
	for x != 0 {
		x = x & (x - 1)
		count++
	}
	return count
}

package main

import "strconv"

func main() {
	y, err := strconv.ParseInt("100", 16, 64)
	println(y,err)
}

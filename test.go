package main

import "fmt"

func main() {

	const name = 1 << iota
	fmt.Println(name)
	fmt.Println(1 << 25)
}

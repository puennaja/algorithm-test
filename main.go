package main

import "fmt"

func a() bool {
	fmt.Println("a")
	return true
}

func b() bool {
	fmt.Println("b")
	return true
}

func main() {
	if a() || b() {
		fmt.Println("done")
	}
}

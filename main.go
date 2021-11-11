package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello you")

	// use feature A
	featureA("mango")

	// use feature B
	fmt.Println(featureB(2, 3))
}

func featureA(str string) {
	fmt.Println(str)
}

func featureB(a, b int64) int64 {
	return a + b
}

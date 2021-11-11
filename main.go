package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello you")

	// use feature A
	featureA("mango")
}

func featureA(str string) {
	fmt.Println(str)
}

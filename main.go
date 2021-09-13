package main //compile project

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
    return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	fmt.Println("multiply", multiply(2, 2))
	totalLength, _ := lenAndUpper("nomankey")
	totalLength2, upperName := lenAndUpper("alpaca")
	fmt.Println(totalLength)
	fmt.Println(totalLength2, upperName)
	repeatMe("nomankey", "alpaca", "monkey", "banana")
}
package main //compile project

import (
	"fmt"
)

func multiply(a int, b int) int {
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))
}
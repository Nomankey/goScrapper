## What is GO?


### what is fmt package
it's a package for formatting
```
fmt.Println("Hello world") //fmt.print line
```

### functions
go can have multiple returns
```

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
    return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, _ := lenAndUpper("nomankey")
	totalLength2, upperName := lenAndUpper("alpaca")
	fmt.Println(totalLength)
	fmt.Println(totalLength2, upperName)
}
//output
// 8
// 6 ALPACA

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("nomankey", "alpaca", "monkey", "banana")
}
//output
// [nomankey alpaca monkey banana]
```
package main

var a = func() int { return 1 }()
var b = func() int { return 2 }()

func main() {
	sum := add(a, b)
	print(sum)
}

func add(a, b int) int {
	return a + b
}

// go build -o concept
// go tool objdump -s main.main concept
// go tool objdump -s main.add concept

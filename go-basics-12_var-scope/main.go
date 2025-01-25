package main

func main() {
	var panic []int
	recover := append(panic, 1)
	len := len(recover)
	//range := len + 1
	_ = len
}

func min() {}
func max() {}

package closure

import "fmt"

// NewIDGenerator returns a closure that
// generates unique IDs
func NewIDGenerator() func() int {
	n := 1
	return func() int {
		var base = 100
		n++
		return base + n
	}
}

func f() {
	idGen := NewIDGenerator()
	fmt.Println(idGen()) // 102
	fmt.Println(idGen()) // 103
	fmt.Println(idGen()) // 104
}

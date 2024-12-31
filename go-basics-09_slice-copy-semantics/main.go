package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 0
	fmt.Println("s1=", len(s1), s1)
	fmt.Println("s2=", len(s2), s2)

	s3 := s1[:2]
	s3[0] = 1
	fmt.Println("s3=", len(s3), s3)
	fmt.Println("s1=", len(s1), s1)
}

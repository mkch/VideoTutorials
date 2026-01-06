package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	type Member struct {
		ID   int
		Name string
		Age  int
	}

	members := []Member{
		{ID: 3, Name: "Bob", Age: 30},
		{ID: 2, Name: "Alice", Age: 25},
		{ID: 1, Name: "Charlie", Age: 35},
	}

	idMap := make(map[int]*Member)
	for i := range members {
		idMap[members[i].ID] = &members[i]
	}

	slices.SortFunc(members,
		func(m1, m2 Member) int {
			return strings.Compare(m1.Name, m2.Name)
		},
	)

	fmt.Println(idMap[2]) // Will output &{3 Bob 30}
}

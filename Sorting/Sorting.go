package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	strs2 := []string{"Лужнов", "Алексей"}
	sort.Strings(strs2)
	fmt.Println("Strings:", strs2)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:", strs)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

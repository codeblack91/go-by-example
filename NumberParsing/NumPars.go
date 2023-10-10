package main

import (
	"fmt"
	"strconv"
)

func main() {

	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 32)
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	k, _ := strconv.Atoi("10000")
	fmt.Println(k)
	t := fmt.Sprintf("%T", k)
	fmt.Println(t)

	testConv := fmt.Sprintf("%d", 120)
	fmt.Printf("%T", testConv)
	fmt.Println(testConv)

}

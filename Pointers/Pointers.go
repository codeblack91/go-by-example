package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func name(name *string) {
	*name = "Vadim"
}

func main() {
	test := "Lena"
	name(&test)
	fmt.Println(test)
}

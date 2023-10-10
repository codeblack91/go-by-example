package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(err())
}

func err() error {
	return errors.New("test")
}

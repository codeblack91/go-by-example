package main

import "fmt"

func main() {
	message := make(chan string, 3)

	message <- "buffered"
	message <- "channel"
	message <- "test"

	fmt.Println(<-message)
	fmt.Println(<-message)

	fmt.Println(<-message)
}

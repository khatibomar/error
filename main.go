package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("kanna")

	err = Inject(err)
	fmt.Println(err)

	file, line, originalErr := Extract(err)
	fmt.Println(file, line, originalErr)
}

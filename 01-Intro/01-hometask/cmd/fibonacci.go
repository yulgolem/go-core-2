package main

import (
	"fmt"
	"go-core-2/01-Intro/01-hometask/pkg/fib"
)

func main() {
	var x, err = fib.Count(100)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(x)

}

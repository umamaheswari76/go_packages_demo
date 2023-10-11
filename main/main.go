package main

import (
	"fmt"

	"github.com/umamaheswari76/go_packages_demo/calc"
)

func main() {
	res := calc.Add(1, 2)
	fmt.Println("Addition is done using my own package and the result is ",res)
}

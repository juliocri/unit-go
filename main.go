package main

import "fmt"
import "github.com/juliocri/unit-go/sample"

func main() {
	res, _ := sample.CidrToMask(10)
	fmt.Printf("Hapy sampling!\n")
	fmt.Printf("%v\n", res)
}

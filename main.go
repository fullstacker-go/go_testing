package main

import (
	"fmt"
	"go_testing/bench_demo"
)

//go:noinline
func main() {

	emp := bench_demo.ChangeName(2)
	fmt.Println(emp)
}

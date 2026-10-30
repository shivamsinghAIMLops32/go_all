package main

import "fmt"

const (
	Execute = 1 << 0
	Write   = 1 << 1
	Read    = 1 << 2
)

func main() {
	// and operator
	fmt.Println(1 & 1)

	// or is used to combine results
	fmt.Println(1 | 2)

	// permissioned for both read and execute
	permisson := Execute | Read
	fmt.Println(permisson)

	//  xor ^ means exclusive or => if different then 1
	fmt.Println(5 &^ 3)

	oddOREven(4)

	fmt.Println(setBitAtIndex(5, 1))
}

func oddOREven(num int) {
	if num < 0 {
		return
	}

	if num&1 == 1 {
		fmt.Println("odd")
	} else {
		fmt.Println("even")
	}
}

func setBitAtIndex(num, idx int) int {
	return num | 1<<idx
}

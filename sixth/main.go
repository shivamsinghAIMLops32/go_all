package main

import "fmt"

const (
	admin = "admin"
)

func main() {
	// operators

	const s = 10 + 14

	k := 20 / 10
	b := 20 % 2

	fmt.Println(s, k, b)

	// comparison
	isGreater := 1 > 2
	fmt.Println(isGreater)
	isEqual := "hello" == "hello"

	fmt.Println(isEqual)
	isAdmin := "admin" != admin
	fmt.Println(isAdmin)

	// unary
	a := 10
	b = a + 1

	fmt.Println(b)

	// y := a--    // here first y gets value fo a and then a is decremented to 9
	// fmt.Println(y) is invalid in go to assign , in go we can only use i++ ,no prefix unary is supported
}

package main

import "fmt"

func main() {
	// and
	a := 5 & 6 // 101 and 110 => 100 which bcmes 4
	fmt.Println(a)

	// we can use to check num & 1 to check if last bit of a number is set or not

	// OR
	b := 10 | 2 // 1010 or 0010 => 1010  which is 10
	fmt.Println(b)

	c := 10 ^ 2 // xor 1010 xor 0010 => 1000 differnt bit 1 and same bceoms 0
	fmt.Println(c)

	d := 10 << 2 // 1010 becomes 101000 which is 32+8 =40
	fmt.Println(d)

	e := 10 >> 2 // 1010 becomes 10 which is 2
	fmt.Println(e)

	// logical and or and
	// z := ((a>b)&&(c>d))? 5 :10  go doesn't supports TERNARY OPERATOR

	if (a > b) && (c > d) {
		fmt.Println("yey boii")
	} else {
		fmt.Println("oh no")
	}

	if 5+10 == 11 || 10-5 == 5 {
		fmt.Println("hmmmm")
	} else {
		fmt.Println("oh yeah baby")
	}
	if 5+10 <= 11 || 10-5 >= 5 {
		fmt.Println("hmmmm")
	} else {
		fmt.Println("oh yeah baby")
	}
}

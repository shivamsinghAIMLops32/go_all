package main

import "fmt"

func main() {
	// there is bool type
	var isAdmin bool = true
	fmt.Println(isAdmin)

	// isAdmin =  0     we can't use 0 and 1 to assign into bool variable cause they are part of integers
	isAdmin = false

	fmt.Println(isAdmin)

	// float data type , also float stores fraction values too
	price := 20.42923
	fmt.Printf("price of tea is %.2f\n", price) // gonna print till only 2 digits after decimal

	// complex number type  cn uses complex64 as float32 and for complex128 as float64 to represnt cn
	var cn complex64 = complex(20.32, 79.42143)

	fmt.Println(cn)
	fmt.Println(real(cn))
	fmt.Println(imag(cn))
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// GO is static typed language that mean during compile it wants to know the type of data we gonna store
	// there are 3 primitive types -> int,string,bool,and float

	// signed integer

	var a int32 // here a is an integer of 32 bits means 4 bytes as 8 bits = 1 byte
	a = 5       // as its signed int that means we can store negative too so from 2^-31 to 2^31 is the range

	// unsinged integer
	// the range is 0 to 2^32 cause now last bit is free rather than getting fixed represent sign
	var b uint64
	b = 25

	fmt.Printf("a is %d with virtual mem add %p and type is %T and size is %d\n", a, &a, a, unsafe.Sizeof(a))
	fmt.Printf("b is %d with virtual mem add %p and type is %T and size is %d\n", b, &b, b, unsafe.Sizeof(b))

	// string type so used by "string"
	var name string = "jason"
	fmt.Printf("%s\n", name)

	// bytes so in go byte is int8 which used for ascii and then we have int32 which is rune
	// byte is like a character 'char'
	var cName []byte = []byte(name) // so we are doing type conversion from string to []byte
	// so => jason = []byte{106,97,115,111,110} as 106 is j in ascii and so forth
	fmt.Println(cName)

	fmt.Println(string([]byte{66, 82, 89, 89, 65, 78}))
}

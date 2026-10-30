package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// array in go are fixed size we cant append anything to it and it is passed by value (copied everything)
	arr := [...]int{2, 4, 6, 10, 13, 15}

	// THERE IS NO HEADER IN CASE OF ARRAY IN GO SO &ARR[0] == &ARR
	fmt.Println(arr)

	fmt.Println(arr[3])

	arr[4] = 50

	// Changes the index 4
	fmt.Println(arr)

	// The memory address of the array &oArr and the memory address of the first element &oArr[0] point to the exact same byte in memory

	brr := arr // arr and brr both has diffrent array in memory => values were copied from arr to brr
	fmt.Printf("arr has address %p and arr's headers address is %p means that's first values address\n", &arr, &arr[0])
	fmt.Printf("brr has address %p and brr's headers address is %p means that's first values address\n", &brr, &brr[0])

	// string
	// string is also  like array in go but we cant even manipulate them in go even inplace cause header has a fix address
	name := "brock"
	name2 := name

	// so string in go has 2 things in struct on header and another is length
	// so name and name2 inside header points to same underlying array,which is read only
	fmt.Printf("name header has address %p and  name2 header has address %p\n", unsafe.StringData(name), unsafe.StringData(name2))

	// here name gets a new array adrress while name2 has the same underlying address
	name = name + " lesnar"

	fmt.Printf("name header has address %p and  name2 header has address %p\n", unsafe.StringData(name), unsafe.StringData(name2))

	b := string(name[1]) // so basically name => ['b','r','o']
	// we can say that each index of string is a byte stored in an immutable array
	fmt.Println(b)

	em := [...]int{1, 2, 3, 4}
	fmt.Printf("lenght of em is %d and type is %T", len(em), len(em))

	// means create a 4 length array and 0 idx has 1 , and 2nd idx to 10 and 3rd to 20 and rest idx would have zero value of type of array it is holdiing
	k := [4]int{1, 2: 10, 20}
	fmt.Println(k)

	// lets try to change a array value at a particular idx
	oArr := [4]int{1, 2, 3, 4}
	changeOriginal(&oArr)
}

// in go [4]int is different type than [5]int
// so a pointer to *[4]int should be passed as same type *[4]int means address of arr of 4 length
// we already told compiler we are giving you a adress of array whole size is 4 and storing int values
func changeOriginal(arr *[4]int) { // so at runtime compiler should be knowing the length of array we gonna compute for
	for i := range arr {
		arr[i] = i * 2 // go automatically resolved deferecing for index based operation
	}
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// string
	name := "jason"
	fmt.Println(name)
	fmt.Printf("string s has length %d", len(name)) // string doesn't have capcacity

	name = string([]byte{80, 90, 111})

	fmt.Println(name)

	name1 := "brock"
	name2 := name1

	// both string aren't at same address they are 2 differnt string
	fmt.Printf("%s has address %v and %s has address %v\n", name1, &name1, name2, &name2)

	// strings are only read only in go and it always passes as value
	// strings cant be mutated in place like s[0] = 11
	// we cant get index, hack is to convert a string into []byte and then work around as now its a slice

	address := "abc street2"
	arrAdd := []byte(address)

	fmt.Println("before updating slice", address)
	fmt.Println("before updating main slice", arrAdd)
	arrAdd[2] = 'Z'
	arrAdd[len(arrAdd)-1] = '5'
	fmt.Println("after updating slice", address)
	fmt.Println(arrAdd)
	fmt.Println(string(arrAdd))

	// to change a string we can do just replace and concate with other string
	email := "abc@gmail"
	email = "1" + email
	fmt.Println(email)

	// we can do inplace updates too
	location := "new york"
	ll := []byte(location)
	ll[0] = 'N' // or use ascii of 	N
	ll = append(ll, []byte(" city")...)

	location = string(ll)
	println(location)
	// intersting way
	// as string return only a readonly header for location annd a length so we do something cooll with unsafe
	// but here we are commiting  that we won't be changing the ll slice ever cause now a string newLL points to the same header add of ll so ya changing an string isn't allowed
	newLL := unsafe.String(&ll[0], len(ll))
	fmt.Println(newLL)
	fmt.Printf("the address of ll internal array header is %p", &ll[0])
	fmt.Printf("the address of location header is %p", unsafe.StringData(newLL))
}

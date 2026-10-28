package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// string and bytes and runes

	var name string = "brock"
	var city string = string([]byte{78, 89, 67})
	var weather string = string([]byte{'c', 'o', 'l', 'd'}) // these each '' are rune but can be represented using uint8 which is byte it will become 67,..,..,..etc

	fmt.Println(name)
	fmt.Println(city)

	fmt.Println(weather)
	//  bytes in go are uint 8 so anything which can be represented by 8 bits can fit into byte
	var email []uint8 = []byte("jason@gmail.com") // uint 8 and byte is same thing
	fmt.Println(email[3])                         // o is 111
	fmt.Println(string(email[5]))                 // email[5] = 64 which is @

	// error code if we have to store a character or string which have utf8 character like chinese,emoji we can't store in byte
	// cause they some do take 4 byte so range of rune is 1 to 4 bytes for utf8

	// var smile byte = '😂' // as smiley takes 4 bytes if we check the size using unsafe
	// fmt.Println(smile)

	fmt.Println(unsafe.Sizeof('😂'))

	// rune int32
	var cry uint32 = '😭'
	var wine rune = '🍷'
	fmt.Println(cry)
	fmt.Println(wine)
}

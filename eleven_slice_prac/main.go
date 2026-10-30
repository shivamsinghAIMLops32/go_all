package main

import "fmt"

func main() {
	// slices in go are mutable
	sl := []int{} // this is a empty slice
	var sl2 []int // this is a nil slice means it isn't intialized yet unless a value is added

	// both are of 0 length and capacity but sl2 is nil slice not empty
	fmt.Printf("the length of slice sl is %d and capacity is %d\n", len(sl), cap(sl))
	fmt.Printf("the length of slice sl2 is %d and capacity is %d\n", len(sl2), cap(sl2))

	sl = append(sl, 10, 20)
	sl2 = append(sl2, 10, 20)
	// both has different address cause header is differnt
	fmt.Printf("the address of sl is %p and the adress of sl2 is %p\n", &sl[0], &sl2[0])

	sl2 = append(sl2, sl...)
	fmt.Println(sl2)

	// new working
	a := make([]int, 0, 5)
	b := a
	c := b
	fmt.Printf("the len and capacity of a is %d and %d\n", len(a), cap(a))
	fmt.Printf("the len and capacity of b is %d and %d\n", len(b), cap(b))
	fmt.Printf("the len and capacity of c is %d and %d\n", len(c), cap(c))

	c = append(c, 10, 20, 30)
	// even tho a, b, c are pointing to same underlying array
	// but only c header gets a new length
	// cause append returns a updated header
	fmt.Printf("the len and capacity of a is %d and %d\n", len(a), cap(a))
	fmt.Printf("the len and capacity of b is %d and %d\n", len(b), cap(b))
	fmt.Printf("the len and capacity of c is %d and %d\n", len(c), cap(c))
	// fmt.Println(a[2]) we cant access this cause it used length as header of a and b is still poiting as length which is zero
	fmt.Println(a[:3]) // [10 20 30]
	fmt.Println(b[:3]) // [10 20 30]
	fmt.Println(c)     // [10 20 30]

	// what happens is that when we copied b from a and c from b every value of a get copied to b and from b it for copied to c
	// thats why appending to c returns a updtad header the underlying is same but now c has length 3 and a nd b len =0
	a = append(a, 100)

	// here a has length 0 so appending to a returns the new updated header where 100 is written at 0 index so a length now is 1 and len(c) is 3 and b is 0
	// but it will reflect in a and b and c when reslicing which gives a new view to same underlying array
	// appends always put value acc to slices length not capacity
	fmt.Printf("a has len %d and cap %d and header address is %p\n", len(a), cap(a), &a[0])
	fmt.Printf("b has len %d and cap %d and header address is %p\n", len(b[:3]), cap(b[:3]), &b[:3][0])
	fmt.Printf("c has len %d and cap %d and header address is %p\n", len(c), cap(c), &c[0])

	// IMPORTANT
	x := make([]int, 2, 5) // so [0,0,0,0,0] is backing array with 0 value
	// and x has len 2 and capacity 5 and header [0,0] cause len is 2
	x[1] = 2
	// now x has len 2 and capacity 5 and header [0,2]
	x = x[:4]

	// x has length 4 and capacity 5 and array x is [0 2 0 0]
	fmt.Printf("x has length %d and capacity %d and array x is %v", len(x), cap(x), x)
}

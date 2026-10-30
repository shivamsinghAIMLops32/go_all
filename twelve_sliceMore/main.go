package main

import "fmt"

func main() {
	// play with slices
	s := make([]int, 0, 5)
	b := s

	b = append(b, 10, 20, 30)
	fmt.Printf("b slice has cap %d and len %d\n", cap(b), len(b))
	fmt.Printf("s slice  has cap %d and len %d\n", cap(s), len(s))

	// b slice has cap 5 and len 3
	// s slice  has cap 5 and len 0
	// b slice [100 200 300] has cap 5 and len 3
	// s slice [100 200 300 400 500] has cap 5 and len 5
	s = append(s, 100, 200, 300, 400, 500)
	fmt.Printf("b slice %v has cap %d and len %d\n", b, cap(b), len(b))
	fmt.Printf("s slice %v has cap %d and len %d\n", s, cap(s), len(s))

	s[4] = 20
	fmt.Printf("b slice %v has cap %d and len %d\n", b, cap(b), len(b))
	fmt.Printf("s slice %v has cap %d and len %d\n", s, cap(s), len(s))

	s = append(s, 1000)
	fmt.Printf("addres of s header is %p\n", &s[0])
	fmt.Printf("addres of b header is %p\n", &b[0])

	b[1] = 2000
	fmt.Printf("b slice %v has cap %d and len %d\n", b, cap(b), len(b))
	fmt.Printf("s slice %v has cap %d and len %d\n", s, cap(s), len(s))

	k := make([]int, 0, 5)
	k = append(k, 10, 20, 30)
	c := k[1:3]
	fmt.Printf("c slice %v has cap %d and len %d\n", c, cap(c), len(c))
	fmt.Printf("k slice %v has cap %d and len %d\n", k, cap(k), len(k))

	c = append(c[:3], 200, 300)
	fmt.Printf("c slice %v has cap %d and len %d\n", c, cap(c), len(c))
	fmt.Printf("k slice %v has cap %d and len %d\n", k, cap(k), len(k))

	t := make([]int, 3, 5)
	d := t[:2:5]
	d = append(d, 200)
	fmt.Println(d)
}

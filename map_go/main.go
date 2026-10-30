package main

import (
	"fmt"
)

type (
	score map[string]int
)

func main() {
	// nil map
	var o map[int]string
	if o == nil {
		fmt.Println("o is nil bro")
	}
	// score map[type]type
	mc5Score := score{
		"jason": 1,
		"randy": 2,
	}

	codmScore := score{
		"jason": 10,
		"randy": 3,
	}

	fmt.Println(mc5Score["jason"])

	gameIDuserScore := map[int]score{
		1: mc5Score,
		2: codmScore,
	}
	fmt.Println(gameIDuserScore[2]["jason"])

	// we have to intitalized map always cause it will throw PANIC if access without intitalized
	mp := make(map[int]string, 10)
	mp[1] = "roj"
	mp[3] = "brock"
	fmt.Println(mp[3])

	mk := mp
	mk[4] = "john"
	fmt.Println(mk[4])
	mk[3] = "chris"
	fmt.Println(mp[3])
	fmt.Println(mk[3])

	// as map doesnt have anything like capacity so doesnt matter if i change length of any of those map the link between them will always be there
	// and changing one will reflect on other too

	// we can check if  a key exists in map or not
	val, ok := mp[4]
	if !ok {
		return
	}
	fmt.Println(val)
}

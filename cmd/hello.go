package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hi := "hello"
	fmt.Printf("%s go\n", hi)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Int31n(5))
}

// Hello ...
func Hello() string {
	return "hello"
}

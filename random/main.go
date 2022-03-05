package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().Unix()
	rand.Seed(t)
	s := rand.Intn(10)

	println(s)
}

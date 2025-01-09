package main

import "log"

func EvenOrOdd(n int) string {
	log.Println("EvenOrOdd called with", n)
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}

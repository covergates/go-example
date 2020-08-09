package main

import "log"

func sum(a, b int) int {
	return a + b
}

func main() {
	log.Println(sum(1, 2))
}

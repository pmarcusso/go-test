package main

import (
	"github.com/pmarcusso/go-test/calc"
	"log"
)

func main() {

	log.Println("Starting program...")

	arrayList := []int{10, 2, 5, 3, 7}

	calc.OrderList(arrayList)
}

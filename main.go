package main

import (
	"fmt"
)

type PrintFunc func(string)

func Print (s string) {
	fmt.Println(s)
}

func PrintEx(fn PrintFunc) PrintFunc {
	return func(s string) {
		fmt.Println("start")
		fn(s)
		fmt.Println("finish")
	}
}


func main() {
	Print("hello")
	PrintEx(Print)("world")
}

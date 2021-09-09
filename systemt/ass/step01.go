package main

import "fmt"

//go tool compile -S -N -l step01.go
func main() {
	i := add(1, 2)
	fmt.Println(i)
	i2 := add(1, 2)
	fmt.Println(i2)
	//add01(1, 2)
}

func add(a, b float32) float32 {
	fmt.Println(a)

	return a + b
}

//
//func add01(a, b int) int {
//	return add02(a, b)
//}
//
//func add02(a, b int) int {
//	return add(a, b)
//}

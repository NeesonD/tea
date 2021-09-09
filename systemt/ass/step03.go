package main

func main() {
	add4(1, 2)
}

//go:noinline
func add4(a, b int) int {
	c := 3
	d := a + b + c
	return d
}

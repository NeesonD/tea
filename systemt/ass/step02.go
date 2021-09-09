package main

//  go tool compile -S -N -l step02.go
func main() {
	p := Point{2, 5}
	p.VIncr(10)
	p.PIncr(10)
}

type Point struct {
	X int
	Y int
}

func (p Point) VIncr(factor int) {
	p.X += factor
	p.Y += factor
}

func (p *Point) PIncr(factor int) {
	p.X += factor
	p.Y += factor
}

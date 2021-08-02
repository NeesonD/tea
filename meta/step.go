package main

var step1Code = `package main  
  func main() {println("hello go")
}`

const (
	ExampleText = `package main

import "fmt"

type {{.Biz}}Model struct {
	Id   int
	Name string
}

func (e *{{.Biz}}Model) Add()  {

}


func (e *{{.Biz}}Model) Update() {

}

func (e *{{.Biz}}Model) Delete() {

}

func (e *{{.Biz}}Model) One() {
	fmt.Println()
}

func GetTableName()  {
	
}
`
)

// testprogram.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// curl "localhost:8080/hello?name=neeson%20Russia%20mark&year=2020"
// bpftrace -e `uprobe:/home/neeson/go/0908/step01:main.* {printf("%s - %s\n",comm,func);}`
func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	name := query.Get("name")
	year_, _ := strconv.ParseUint(query.Get("year"), 10, 32)
	year := int(year_)
	status := checkSite()
	answer := prepareAnswer(name, year, status)
	writer.Write([]byte(answer + "\n"))
	return
}

//go:noinline
func checkSite() int {
	resultChan := make(chan int)
	go func() {
		resp, err := http.Get("https://www.gophercon-russia.ru")
		if err != nil {
			log.Fatalf("http get failed: %s\n", err)
		}
		resultChan <- resp.StatusCode
	}()

	return <-resultChan
}

//go:noinline
func prepareAnswer(name string, year int, status int) string {
	answer := fmt.Sprintf("Hello, %s %d! Website returned status %d.", name, year, status)
	return answer
}

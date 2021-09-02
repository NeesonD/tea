// client/client.go
package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:7788")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 1000; i++ {
		args := &Args{A: 10, B: 100}
		var reply int

		// <-- 目的就是为了记录 Call 的耗时
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
	}
}

// go build -gcflags="-N -l" -o client client.go

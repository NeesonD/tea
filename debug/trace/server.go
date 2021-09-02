// server/server.go
package main

import (
	"log"
	"net"
	"net/rpc"
	"time"
)

type Args2 struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args2, reply *int) error {
	time.Sleep(3000 * time.Millisecond) // 假设每次调用服务端需要处理 3s
	*reply = args.A * args.B
	return nil
}

func main() {
	server := rpc.NewServer()
	server.Register(new(Arith))

	l, err := net.Listen("tcp", "127.0.0.1:7788")
	if err != nil {
		log.Fatalf("failed to connect address: [%s], error: %v", "127.0.0.1:7788", err)
	}

	for {
		conn, err := l.Accept()

		log.Printf("accept new connection: %v\n", conn.RemoteAddr().String())

		if err != nil {
			log.Println("listener Accept error")
			time.Sleep(100 * time.Millisecond)
			continue
		}

		go server.ServeConn(conn)
	}
}

// go build -gcflags="-N -l" -o server server.go

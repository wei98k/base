package interview

import (
	"fmt"
	"log"
	"net"
	"testing"
	"time"
)

func TestMyFib(t *testing.T) {
	go MySpinner(100 * time.Microsecond)
	const n = 45
	fibN := MyFib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func TestHandleConn(t *testing.T) {
	listtener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listtener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		HandleConn(conn)
	}
}

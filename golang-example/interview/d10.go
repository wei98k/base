package interview

import (
	"fmt"
	"io"
	"net"
	"time"
)

func MySpinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func MyFib(x int) int {
	if x < 2 {
		return x
	}
	return MyFib(x-1) + MyFib(x-2)
}

func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

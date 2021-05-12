package ch07

import (
    "testing"
)

func TestWrite(t *testing.T) {
    var c ByteCounter
    strLen, _ := c.Write([]byte("hello"))

    
    if strLen != 5 {
       t.Fatalf("expect 5, return %d", strLen) 
    }
    
}

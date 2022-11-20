package main

import(
    "fmt"
)

type T struct {
    a int
    b float32
    c string
}

func main() {
     t := &T{7, -2.35, "abc\tdef"}
     r := t.String()
     fmt.Println(r)
}

func (t *T) String() string {
    return fmt.Sprintf("%d / %f %q", t.a, t.b, t.c)
}

// 输出: 7 / -2.350000 "abc\tdef"

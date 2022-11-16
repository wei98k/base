package main

import(
    "fmt"
)

var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
)

// 键值对调
func main() {
    // 新的map 
    m := make(map[int]string, len(barVal))
    for k, v := range barVal {
        m[v] = k
    }
    fmt.Println(m)
}

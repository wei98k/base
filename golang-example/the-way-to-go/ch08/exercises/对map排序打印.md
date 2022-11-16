package main

import(
    "fmt"
    "sort"
)

var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
)

func main() {
    fmt.Println("未排序")
    for k, v := range barVal {
        fmt.Printf("k: %s, v: %d\n", k, v)
    }
    keys := make([]string, len(barVal))
    i := 0
    for k := range barVal {
        keys[i] = k
        i++
    }
    sort.Strings(keys)
    for _, k := range keys {
        fmt.Printf("k: %s, v: %d\n", k, barVal[k])
    }
}

package main

import(
    "fmt"
)

// 定义int别名
type TZ int

// 定义时区
const (
    HOUR TZ = 60 * 60
    UTC TZ = 0 * HOUR
    EST TZ = -5 * HOUR
    CST TZ = -6 * HOUR
)

var timeZones = map[TZ]string{
    UTC: "Universal Greenwich time",
    EST: "Eastern Standard time",
    CST: "Central Standard time"}

func (tz TZ) String() string {
    if zone, ok := timeZones[tz]; ok {
        return zone
    }
    return ""
}

func main() {
    fmt.Println(EST)
    fmt.Println(0 * HOUR)
    fmt.Println(-6 * HOUR)
}

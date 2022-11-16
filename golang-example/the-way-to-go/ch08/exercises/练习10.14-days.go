package main

import(
    "fmt"
)

type Day int

const(
    MO Day = iota
    TU
    WE
    TH
    FR
    SA
    SU
)

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (day Day) String() string {
    return dayName[day]
}
// 思路分析: 
// 1. 定义int别名Day
// 2. 定义常量 iota自增
// 3. 定义一个一周的名字字符串变量
// 4. 初始化赋值-string函数会被fmt自动使用(大部分方法)
func main() {
    var th Day = 3
    fmt.Printf("The 3rd day is: %s\n", th.String())
}

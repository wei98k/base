package main

import(
    "fmt"
    "time"
)
// 定义vcard.go
type VCard struct {
    FirsName string
    LastName string
    NickName string
    BirtDate time.Time
    Photo string
    Addresses map[string]*Address
    
}
// 定义结构体Address
type Address struct {
    Street string // 街道
    HouseNumber uint32 // 住址编号 
    HouseNumberAddOn string
    POBox string
    ZipCode string
    City string
    Country string
}

func main() {
    // 初始化struct 不指定字段名字，就一定要按顺序赋值
    addr1 := &Address{"阳光大道", 89, "", "", "2600", "深圳", "中国"}
    addr2 := &Address{"苏喂街道", 12, "", "", "9610", "样样", "美国"}
    // 初始VCard->map
    addrs := make(map[string]*Address)
    addrs["youth"] = addr1
    addrs["nowh"] = addr2
    birthdt := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local) 
    // birthdt := time.Date(1996, 1, 17, 15, 4, 5, 0, time.Local)
    photo := "MyDocuments/MyPhotos/photo1.jpg"
    vcard := &VCard{"Ivo", "Balbaert", "", birthdt, photo, addrs}
    fmt.Printf("Here is the full VCad: %v\n", vcard)
    fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}

package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	A1 string
	A2 string
}

//
//
//两组题目，选一组做
// A:
// 1. 将切片去重复
// 2. 冒泡排序
// 3. 用接口实现Duck typing

// B:
// 1. 用反射遍历结构体字段，包含嵌套的结构体，并打印字段名。
//

type Atest struct {
	Name string `yaml:"name" name:"username"`
	Sex  int64  `yaml:"sex" name:"sex"`
	Age  int16  `yaml:"age" name:"age"`
}

func main() {
	d := Atest{
		Name: "name",
		Sex:  1,
		Age:  16,
	}
	t := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	for k := 0; k < t.NumField(); k++ {
		fmt.Println("name:", fmt.Sprintf("%+v", t.Field(k).Name), ", value:", fmt.Sprintf("%v", v.Field(k).Interface()))
	}
}

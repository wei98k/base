package main

//利用反射解析json返回struct 同样也可以让struct返回json
import(
	"fmt"
	"reflect"
)

type Goods struct {
	// 自增ID 反引号是标签-可以任意字符
	ID int64 `json:"id"`
	// 标题
	Title string `json:"title"`
	// 价格
	Price float64 `json:"price"`
	// 简介
	Info string `json:"info"`
}

func main() {

	// 给结构体赋值
	g := Goods{ID:1, Title:"IPhone 13 Pro", Price: 33.33, Info:"are you ok ?"}

	// 声明变量t
	var t reflect.Type

	// 参数接收是一个任意类型 返回一个 type
	t = reflect.TypeOf(g)

	// 声明变量v
	var v reflect.Value
	v = reflect.ValueOf(g)

	// 组装json字符串
	json := "{"
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") != "-" {
			json += "\"" + t.Field(i).Tag.Get("json") + "\":\"" + v.FieldByName(t.Field(i).Name).String() + "\""
		}
	}
	json += "}"


	fmt.Println(json)
}
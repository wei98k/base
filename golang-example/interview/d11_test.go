package interview

import (
	"fmt"
	"reflect"
	"testing"
)

type BaseData struct {
	// mysql
	DbUrl      string `yaml:"db_url" name:"数据库地址"`
	DbUser     string `yaml:"db_user" name:"数据库用户名"`
	DbPassWord string `yaml:"db_pass_word" name:"数据库密码"`
	DbName     string `yaml:"db_name" name:"数据库名"`
}

func TestBubSort(t *testing.T) {

	num := []int{1, 5, 3, 9, 2}
	BubSort(num)
	fmt.Println(num)
}

func TestQuickDescSort(t *testing.T) {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	QuickDescSort(arr, 0, len(arr)-1)
	fmt.Println("quickDescendSort: ", arr)
}

func TestQuickAscSort(t *testing.T) {
	var arr = []int{19, 8, 16, 15, 23, 34, 6, 3, 1, 0, 2, 9, 7}
	QuickAscSort(arr, 0, len(arr)-1)
	fmt.Println("quickAscSort: ", arr)
}

func TestRemoveDuplicationMap(t *testing.T) {
	var arr = []string{"a", "d", "a", "k", "o", "d"}
	newArr := RemoveDuplicationMap(arr)
	fmt.Println(newArr)
}

func TestRemoveDuplicationSort(t *testing.T) {
	var arr = []string{"a", "d", "a", "k", "o", "d"}
	newArr := RemoveDuplicationSort(arr)
	fmt.Println(newArr)
}

func TestForStruct(t *testing.T) {
	d := BaseData{
		DbUrl:      "url",
		DbUser:     "user",
		DbPassWord: "pw",
		DbName:     "name",
	}
	t1 := reflect.TypeOf(d)
	v := reflect.ValueOf(d)
	for k := 0; k < t1.NumField(); k++ {
		fmt.Println("name:", fmt.Sprintf("%+v", t1.Field(k).Name), ", value:", fmt.Sprintf("%v", v.Field(k).Interface()), ", yaml:", t1.Field(k).Tag.Get("yaml"))
	}
}

func TestMaxSlidingWindow2(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	newNum := MaxSlidingWindow2(nums, 3)
	fmt.Println(newNum)
}

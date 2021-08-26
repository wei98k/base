package main

import (
	"example/util"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 1 {
		fmt.Println(args[1])
	}

	// 返回内核提供的主机名
	hostname, err := os.Hostname()
	fmt.Println(hostname, err)

	dirContext, err1 := os.ReadDir("/Users/jw/workspace/base/golang-example/eg-os/")

	fmt.Println(err1)

	for k, v := range dirContext {
		fmt.Println(k, v.Name())
	}

	println("方式1 os.arg: ", args[0])

	pwd, err := os.Getwd()

	println("方式2 Getwd: ", pwd, err)

	println("方式3 getCurrentAbPathByExecutable: ", util.GetCurrentAbPathByExecutable())

	println("方式4 getCurrentAbPathByCaller: ", util.GetCurrentAbPathByCaller())

	println("getCurrentAbPath: ", util.GetCurrentAbPath())

	//===== 改变当前目录
	// os.Chdir("../")

	// pwd1, err := os.Getwd()

	// println("方式2 Getwd: ", pwd1, err)

	//===== 设置文件权限

	err2 := os.Chmod("./tmp/a.txt", 0644)

	fmt.Println(err2)

	MyCreateTemp()
}

func MyCreateTemp() {

	// 创建一个临时文件
	f, err := os.CreateTemp("", "example")
	if err != nil {
		log.Fatal(err)
	}
	// 程序执行完 删除文件
	defer os.Remove(f.Name())
	// 写入临时文件
	if _, err := f.Write([]byte("content")); err != nil {
		log.Fatal(err)
	}
	// 关闭临时文件
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
	println("run MyCreateTemp")
}

package main

import (
	"example/util"
	"fmt"
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

}

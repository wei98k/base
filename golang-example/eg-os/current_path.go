package egos

import (
	"example/util"
	"os"
)

func GetCurrentAbPath() {

	args := os.Args

	println("方式1 os.arg: ", args[0])

	pwd, err := os.Getwd()

	println("方式2 Getwd: ", pwd, err)

	println("方式3 getCurrentAbPathByExecutable: ", util.GetCurrentAbPathByExecutable())

	println("方式4 getCurrentAbPathByCaller: ", util.GetCurrentAbPathByCaller())

	println("getCurrentAbPath: ", util.GetCurrentAbPath())

}

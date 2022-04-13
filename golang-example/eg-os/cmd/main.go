package main

import (
	"bytes"
	"errors"
	"example/util"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	// args := os.Args

	// if len(args) > 1 {
	// 	fmt.Println(args[1])
	// }

	// 返回内核提供的主机名
	// hostname, err := os.Hostname()
	// fmt.Println(hostname, err)

	// dirContext, err1 := os.ReadDir("/Users/jw/workspace/base/golang-example/eg-os/")

	// fmt.Println(err1)

	// for k, v := range dirContext {
	// 	fmt.Println(k, v.Name())
	// }

	// println("方式1 os.arg: ", args[0])

	// pwd, err := os.Getwd()

	// println("方式2 Getwd: ", pwd, err)

	// println("方式3 getCurrentAbPathByExecutable: ", util.GetCurrentAbPathByExecutable())

	// println("方式4 getCurrentAbPathByCaller: ", util.GetCurrentAbPathByCaller())

	// println("getCurrentAbPath: ", util.GetCurrentAbPath())

	//===== 改变当前目录
	// os.Chdir("../")

	// pwd1, err := os.Getwd()

	// println("方式2 Getwd: ", pwd1, err)

	//===== 设置文件权限

	// err2 := os.Chmod("./tmp/a.txt", 0644)

	// fmt.Println(err2)

	// MyCreateTemp()

	//===== 读取文件

	// 	fmt.Println("GetProjectRoot: ", util.GetProjectRoot())

	// MAINFOR:
	// 	for {
	// 		fmt.Println("")
	// 		fmt.Println("*******请选择示例：*********")
	// 		fmt.Println("1 表示 io.Reader 示例")
	// 		fmt.Println("2 表示 io.ByteReader/ByteWriter 示例")
	// 		fmt.Println("q 退出")
	// 		fmt.Println("***********************************")

	// 		var ch string
	// 		fmt.Scanln(&ch)

	// 		switch ch {
	// 		case "1":
	// 			ReaderExample()
	// 		case "2":
	// 			ByteRWerExample()
	// 		case "q":
	// 			fmt.Println("程序退出")
	// 			break MAINFOR
	// 		default:
	// 			fmt.Println("输入错误")
	// 			continue
	// 		}
	// 	}

	//===== 创建文件
	WriteFile()

	//===== 文字输入到标准输出
	// reader := bytes.NewReader([]byte("golang is ok"))
	// reader.WriteTo(os.Stdout)
	// fmt.Println("")

	//===== pipeReader 和 PipeWriter 类型
	// pipeReader, pipeWriter := io.Pipe()
	// go PipeWrite(pipeWriter)
	// go PipRead(pipeReader)
	// time.Sleep(30 * time.Second)

	//===== 创建目录示例
	// myMkdir()
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

func ReaderExample() {
FOREND:
	for {
		readerMenu()
		var ch string
		fmt.Scanln(&ch)
		var (
			data []byte
			err  error
		)
		switch strings.ToLower(ch) {
		case "1":
			fmt.Println("输入不多于9个字符, 以回车结束: ")
			data, err = ReadFrom(os.Stdin, 11)
		case "2":
			file, err := os.Open(util.GetProjectRoot() + "/01.txt")
			if err != nil {
				fmt.Println("打开文件错误: ", err)
				continue
			}
			data, err = ReadFrom(file, 9)
			file.Close()
		case "3":
			data, err = ReadFrom(strings.NewReader("from string"), 12)
		case "4":
			fmt.Println("未实现")
		case "b":
			fmt.Println("返回上级菜单")
			break FOREND
		case "q":
			fmt.Println("程序退出")
			os.Exit(0)
		default:
			fmt.Println("输入错误")
			continue
		}

		if err != nil {
			fmt.Println("数据读取失败, 可以试试从其他输入源读取")
		} else {
			fmt.Printf("读取到的数据是: %s\n", data)
		}
	}

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func readerMenu() {
	fmt.Println("")
	fmt.Println("*******从不同来源读取数据*********")
	fmt.Println("*******请选择数据源，请输入：*********")
	fmt.Println("1 表示 标准输入")
	fmt.Println("2 表示 普通文件")
	fmt.Println("3 表示 从字符串")
	fmt.Println("4 表示 从网络")
	fmt.Println("b 返回上级菜单")
	fmt.Println("q 退出")
	fmt.Println("***********************************")
}

func ByteRWerExample() {
FOREND:
	for {
		fmt.Println("请输入要通过WriteByte写入的一个ASCII字符 (b: 返回上级; q: 退出): ")
		var ch byte
		fmt.Scanf("%c\n", &ch)
		switch ch {
		case 'b':
			fmt.Println("返回上级菜单")
			break FOREND
		case 'q':
			fmt.Println("程序退出")
			os.Exit(0)
		default:
			buffer := new(bytes.Buffer)
			err := buffer.WriteByte(ch)
			if err == nil {
				fmt.Println("写入一个字节成功！准备读取该字节。。。。")
				newCh, _ := buffer.ReadByte()
				fmt.Printf("读取的字节： %c\n", newCh)
			} else {
				fmt.Println("写入错误")
			}
		}
	}
}

func WriteFile() {
	// 相对当前目录创建文件
	// file, err := os.Create("writeAt.txt")
	// 指定目录创建文件
	file, err := os.Create("./mydir/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang very NB-sdfslfkjljklsdfsd")
	n, err := file.WriteAt([]byte("Go Go Go"), 10)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func PipeWrite(writer *io.PipeWriter) {
	data := []byte("Golang very happy")
	for i := 0; i < 3; i++ {
		n, err := writer.Write(data)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("写入字节 %d\n", n)
	}
	writer.CloseWithError(errors.New("写入端已关闭"))
}

func PipRead(reader *io.PipeReader) {
	buf := make([]byte, 128)
	for {
		fmt.Println("接口端开始阻塞5秒钟。。。")
		time.Sleep(5 * time.Second)
		fmt.Println("接收端开始接收")
		n, err := reader.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("接到字节: %d\n buf内容: %s\n", n, buf)
	}
}

func myMkdir() {
	// Mkdir、MkdirAll、MkdirTemp

	// 创建一个777权限的目录
	// 如果文件夹已存在也不会报错的
	os.Mkdir("./mydir", fs.ModePerm)
	// os.Chmod("./mydir", 0777)

	// 创建多层级目录
	// os.MkdirAll("./mydir/a/b", fs.ModePerm)
}

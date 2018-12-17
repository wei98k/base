# go 

[官方网站](https://golang.org/)

## 入坑资料

[视频-入门视频教程](https://www.imooc.com/learn/968)
[视频-Golang爬虫上手指南完整版](https://www.bilibili.com/video/av31551627/)
[文章-GO入门指南](https://go.fdos.me)
[文章-mac上安装go环境](https://blog.csdn.net/xiaoquantouer/article/details/79985650)

## 安装

- mac 10.13.1 17B1003
- go1.11.4 darwin/amd64

安装-更新-卸载

### 安装

1. 下载安装包

https://golang.org/doc/install#macos  (双击安装 下一步 ok) 安装完毕后, 程序文件目录位置: /usr/local/go 

2. 环境变量设置(GOROOT和GOPATH)

我自己的配置如下

` ~/.bash_profile` 如果文件不存在直接vi打开创建并写入

```
GOROOT=/usr/local/go
export GOROOT
export GOPATH=/Users/jw/mycodelife/base/test/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOBIN:$GOROOT/bin
```

`source ~/.bash_profile` 执行命令让环境生效

`go env` 查看go配置信息

`go version` 查看go版本信息

配置说明

- GOROOT： go安装目录
- GOPATH：go工作目录
- GOBIN：go可执行文件目录
- PATH：将go可执行文件加入PATH中，使GO命令与我们编写的GO应用可以全局调用

### 卸载

直接移除GO程序目录即可 `/usr/local/go`

To remove an existing Go installation from your system delete the go directory. This is usually /usr/local/go under Linux, macOS, and FreeBSD or c:\Go under Windows.

## hello world

cd $GOTPAH

touch hello.go 将下面代码粘贴到文件中

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

go build hello.go  相同目录中将出现多一个hello文件

./hello 输出hello, world. 


## 关于IDE编辑器

**vscode**

提示: The "go-outline" command is not available. User "go get -v .... " 点击 install. 会在目录下`/Users/jw/go/bin`下载文件.

## FAQ

T: 2021年12月29日19:39:45
Q: 如何分析当前程序是否存在goroutine泄露问题？
A: 

T: 2021年12月29日10:38:24
Q: 当前配置是否可以在1s内完成127个html的下载？

假设网络状态ping: icmp_seq=0 ttl=50 time=83.394 ms

mac os air 8G
machdep.cpu.core_count: 2
machdep.cpu.thread_count: 4

T: 2021年12月28日21:12:41
Q: 目前怀疑是内存泄漏问题 goroutine
A:

T: 2021年12月28日19:53:09
Q: 是否可以返回当前设置GOMAXPROCS的值? 目前这个函数返回的是上一次设置的值.
A: 

T: 2021年12月28日10:54:28
Q: make可以任何类型 包括自定义类型吗 比如io.Reader？
A: 

T: 2021年12月28日10:55:08
Q: make struct内部元素的如何初始化？
A: 

T: 2021年12月28日10:24:32
Q: 如果统计当前程序运行了多个协程 正在运行的和执行开始到结束的？
A: 

T: 2021年12月28日10:24:16
Q: 把代码改成协程的方式有几种？
A: 

Q: 如何过滤出页面中的链接？
A: 链接都是带a标签的, 把页面当成一个字符串，匹配出全部的a链接

## TODO


- [x] 测试第网上下载程序 2021年12月29日09:52:51 6.41s run time 抓取时间还是太长了



## 版本迭代

- v5.1-main 抓取链接保存到url.txt文件中用于测试网上的抓取函数 https://blog.csdn.net/zhsheng26/article/details/87983149
- v4.5-main 增加协程数量限制
- v4.4-main 改wg传入函数的方式来递减协程
- v4.3-main 增加channel方式但出现错误 savefile err: http: read on closed response body
- v4.1-main 直接在 line 82 HttpGet2 前加go
- v4-main 整理代码 实现下载HTML和资源文件 2021年12月26日
- v1-main 实现基本的抓取下载功能 2021年12月23日

## 执行时间


- v4.5-main 4.22s run time 2021年12月28日19:15:51
- v4.5-main 4.29s run time 2021年12月28日19:13:31

- v4.4-main 6.14s run time 2021年12月28日18:57:17
- v4.4-main 5.74s run time 2021年12月28日18:56:40
- v4.4-main 6.08s run time 2021年12月28日18:55:57

- v4.2-main 7.32s run time 2021年12月28日10:26:49
- v4.2-main 11.10s run time 2021年12月28日10:26:11
- v4.2-main 8.24s run time 2021年12月28日10:23:37

- v4-main 10.21s run time 2021年12月28日09:53:34
- v4-main 9.80s run time  2021年12月28日09:53:06
- v4-main 10.32s run time 2021年12月28日09:52:29

## 涉及知识点

- 爬虫的基本思路
- 文件处理思路
- go-html扩展包的使用
- 递归获取链接
- 切片结构保存链接
- 如何抓取图片、CSS、JS文件？
- os包 创建文件、写入文件、创建目录
- bytes包 Split 字符串分割
- 并发获取URL
- 并发写入文件
- 并发创建文件夹
- html转PDF [github-go-wkhtmltopdf](https://github.com/SebastiaanKlippert/go-wkhtmltopdf)

## 可并发点

- 每个页面遍历查找链接
- 下载资源文件(但是不能无限创建携程 20个携程)
- 下载每个页面(请求页面, 创建目录，保存文件)

## 功能实现

- [ ] 如何在返回的页面找到自己想要的标签，并提取内容
- 抓取指定域名的全部网页
- 将抓取的页面转成PDF
- 可选指定区域

## 条件

只抓取当前网站，遇到外链就忽略掉

## 代码流程

1. 获取初始页面
2. 获取页面全部的链接 外连接除外
3. 把指定页面的区域转成PDF文件
4. 把PDF合成一个文件



## 相关已实现功能的项目

[python-使用asyncio爬取gitbook内容输出pdf](https://juejin.cn/post/6844903799501357070)

[python-github-使用asyncio爬取gitbook内容输出pdf](https://github.com/fuergaosi233/gitbook2pdf/blob/master/README_zh.md)

## 资源抓取

[中文网站排行榜-站长之家](https://top.chinaz.com/alltop/)

## 参考资料

runtime与Goroutine泄漏

[跟面试官聊 Goroutine 泄露的 6 种方法](https://segmentfault.com/a/1190000040161853)

[Golang cpu的使用设置--GOMAXPROCS](https://blog.csdn.net/lanyang123456/article/details/80832929)

[Golang runtime.NumGoroutine函数代码示例](https://vimsky.com/examples/detail/golang-ex-runtime---NumGoroutine-function.html)

[用 runtime 包做 Go 应用的基本监控](https://studygolang.com/articles/14410)

[Go 笔记之如何防止 goroutine 泄露（二）](https://zhuanlan.zhihu.com/p/75555215)

-----

下载页面方式

[goLang爬取html](https://blog.csdn.net/kansas_lh/article/details/104505290)

[Go 语言下载文件 http.Get() 和 io.Copy()](https://www.twle.cn/t/384)

-----

创建文件或文件夹权限问题

[Golang 创建文件权限问题](https://zhuanlan.zhihu.com/p/33692995)

## 参考资料-下载文件-保存文件-下载器

[golang协程并发下载多个文件](https://blog.csdn.net/zhsheng26/article/details/87983149)

[聊一聊Go实现的多协程下载器--gopeed-core](https://juejin.cn/post/6979241352580038692)

[github-gopeed-core](https://github.com/monkeyWie/gopeed-core)

[强大高效而精简易用的Golang爬虫框架Colly](https://segmentfault.com/a/1190000023808304)

[github-colly](https://github.com/gocolly/colly)

[Golang 实现的文件上传下载工具](http://www.codebaoku.com/it-go/it-go-202672.html)





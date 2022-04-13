# SKB-短链接服务

26+26+10=62^5 约等于 9亿多个链接

最终实现的效果: http://s.cn/Yu72k [域名][5个固定字符]

## 版本

- v1.0.0 以最简单的方式实现短链接服务, 无数据库.
    - 访问http://127.0.0.1:1234/{key} 跳转对应的长链接中
- v2.0.0 MySQL存储服务
- v3.0.0 Redis存储服务

## FAQ

- 如何避免或最大程度减少key冲突？
- 如何让URL变得更短呢？
- 如何抗住大qps？大流量？
- url hash如何设计？
- 存储如何设计？
- 生成1w个短链接需要多长时间？不同版本-不同的CPU核心数量
- 如何对短链接服务做压力测试？
- 如何兼并用户可以自定义短链接部分字符？

### 如何对短链接服务做压力测试？

- 测试环境(单核情况、多核情况、单机情况、分布式情况)
- 测试结果

1. 数据准备, 批量生成链接(1w, 10w, 100w)

## 核心功能实现思路

- 直接把URL进行md5
- 随机生成一个字符串 唯一KEY
- ID生成器(10转62进制)
- hash-MurmurHash算法(10转62进制)

## 存储设计

### table skb_url

- id
- short_key
- long_url

## 测试

- 数据准备-1w个URL生成 模拟1w个用户在同一时间内有短链转换需求

## 相关项目

[wangkaiyan/shorturl](https://github.com/wangkaiyan/shorturl)

[jwma/jump-jump](https://github.com/jwma/jump-jump)

## 测试资料

[ab和wrk测试工具](https://www.jianshu.com/p/c3046b5a08ff)

[ab测试URL传参数](https://www.cnblogs.com/taiyonghai/p/5810150.html)

[wrk 性能测试带参数的接口](https://blog.csdn.net/ffyyhh995511/article/details/102517754)


## 相关资料

[github-murmur3](https://github.com/spaolacci/murmur3)

[docs-murmur3](https://pkg.go.dev/github.com/spaolacci/murmur3)

[10转62进制](https://tomoya92.github.io/2017/04/07/golang-hex-conversion/)


[concurrent map writes](https://cloud.tencent.com/developer/article/1821143)








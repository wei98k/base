
# PHP面试题

- GET和POST区别

1. GET参数URL可见, POST通过HTTP机制参数不可见.
2. GET数据量小, POST数据量大(理论上不限制).
3. GET安全性低, post高. 
4. RESTful看 GET查询资源, POST增加资源.

- Cookie 与 Session 区别

1. Session是在服务端保存的一个数据结构, 用来跟踪用户的状态, 这个数据可以保存在集群、数据库、文件中.
2. Cookie是客户端保存用户信息的一种机制, 用来记录用户的一些信息, 也是实现Session的一种方式. 
3. 如果浏览器禁用Cookie, 同时session也会失效(可以通过其他方式实现, 比如url中传递session_id).
4. 维持一个回话核心就是客户端的唯一标识, 即session_id.

[参考资料1](https://www.zhihu.com/question/19786827)


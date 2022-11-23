
- [ ] B树和B+树具体实现？
- [x] 打印文章 <MySQL的B-Tree索引底层结构以及具体实现原理详解>
- [ ] 优化腾讯云每次部署很慢的问题(抽出组建、不让其在云上重新安装)
- [ ] 优化短链接跳转很慢的问题(该成vue方式、以单页面调用)
- [ ] N个集合去重问题(mongodb去重、读出数据后JS去重)

- [ ] var let const的区别？
- [ ] Promise和sync
- [ ] 属性、类、封装
- [ ] Module
- [ ] vue.js
- [ ] 跳转微信小程序H5页面改造成vue
- [ ] 油猴抓取懂车帝、汽车之家、社区数据
- [ ] 模拟100用户、1000用户、1W用户在同一时间点领取优惠券

- [ ] 计算合适的前缀长度(p154 193)
  - [ ] 创建city_demo从city表导入数据
  - [ ] 复制表数据然后使用rand打乱表的数据
  - [ ] 通过公式计算出适合表的前缀索引长度

## 2022年11月20日 周日

### Work-todo

### Work-log

===mac制作ubuntu镜像启动盘

镜像下载: https://developer.aliyun.com/mirror/

Mac下烧一个Ubuntu的启动盘 https://www.baifachuan.com/posts/1e92edd1.html

弹出U盘 `diskutil eject /dev/diskN`

FAQ: 

问题: `dd: /dev/rdisk4: Resource busy` 
处理方案: 卸载子分区有效  https://unix.stackexchange.com/questions/271471/running-dd-why-resource-is-busy

```
/dev/disk2
   #:                       TYPE NAME                    SIZE       IDENTIFIER
   0:     FDisk_partition_scheme                        *8.0 GB     disk2
   1:             Windows_FAT_16 wr_usb_efi              134.2 MB   disk2s1
   2:                      Linux                         1.1 GB     disk2s2
vgsprasad-mbp:~ vgsprasad$ diskutil umount /dev/disk2s1
Volume wr_usb_efi on disk2s1 unmounted
vgsprasad-mbp:~ vgsprasad$ diskutil umount /dev/disk2s2
disk2s2 was already unmounted
```

===如何在 Mac OS X 下制作可以在 PC 上启动的PE系统?

https://www.zhihu.com/question/390946652

===配置家里的rustdesk无法远程控制,需要如何排查问题？

目前机器分别有深圳2台客户端(A B)，惠州1台客户端(C)，云服务器1台(S)

云服务器已部署完成，并且A和B客户端远程连接都是正常，只有C客户端无法远程控制。C客户端的ID和key的配置也都正确，软件下方也提示连接就绪。

尝试:

1. 通过A连接C，提示连接中，最后提示无法连接中继服务器 终止
2. 通过C连接B，提示连接中，最后提示无法连接中继服务器 终止
3. 去掉A和C的连接S的ID和KEY用官方的服务器，连接成功

客户端C的防火墙问题?
杀毒软件的问题?
家里的路由器问题？
客户端C中的rustdesk配置的问题？


### Work-review


## 2022年11月20日 周日

### Work-todo

- 格式化硬盘，重新安装ubuntu系统 带界面的
- ubuntu连接wifi
- 更新系统，安装必要的工具命令。
- 安装docker部署相关程序
- 整理简历、复习面试题


### Work-log

### Work-review

## 2022年11月20日 周日
### Work-todo

### Work-log

重置密码 passwd 出现 Authentication token manipulation error

进入高级模式 https://zhuanlan.zhihu.com/p/145417594

```
mount -o rw,remount /
```

### Work-review


## 2022年11月17日 周四
### Work-todo

- 登录界面回车键无效果
- 编辑或添加后页面数据不加载
- 扩展包如何开发调试、安装 (车库扩展包)
- 界面开发表单字段需要大量的重复的工作(如何减少重复的动作)

### Work-log

===整理larvel文档大纲制作map

===思维导图 markdown 工具 Markmap

Markmap缺点

- 当节点数量多的时候，无法通过可视化编辑 寻找对应的节点就变的不方便了

Makrmap总结: 对于节点数量少的时候还是比较好用的

markmap官网 https://markmap.js.org/

markmap-vscode插件 https://marketplace.visualstudio.com/items?itemName=gera2ld.markmap-vscode

markmap介绍文章: https://juejin.cn/post/7000874049333100551


### Work-review

## 2022年11月16日 周三
### Work-todo

### Work-log

===laravel php

laravel api资源排序
```
BrandResource::collection(Brand::all()->sortByDesc("id"))
```
### Work-review

## 2022年08月14日08:30:42

~~整理时间: 2022年08月14日16:54:37 位置:blog~~

测试/usr/share/dict/words数据进行CRC32计算后冲突率

  - [x] MAC电脑中找到`/usr/share/dict/words`文件, 上传到服务器导入到数据中
  - [x] 文件如何导入到数据库中? 1 通过其他的脚本 2 直接使用mysql导入
  - [x] 查找冲突记录`select word,cc FROM words group by cc`
  - [x] 对比表的大小(索引和数据总大小)
  - [x] 对比查询速度

```
DROP TABLE IF EXISTS `words`;
CREATE TABLE `words2` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `word` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE words ADD COLUMN cc_word int unsigned not null default 0 AFTER word;
```

```
ALTER TABLE `words2` ADD INDEX idx_word ( `word` )
ALTER TABLE `words` ADD INDEX idx_word ( `cc_word` )
```

```
insert into words(word) values (load_file('/home/mysqlapp/word-test.txt'));
insert into words(word) values (load_file('/var/lib/mysql-files/word-test.txt'));
insert into words(word) values ("ddd");

insert into words(word, cc_word) values ("okokok", 2238777324);

load data low_priority infile "/var/lib/mysql-files/words" into table words2(word); 

update words set cc_word = crc32(word);
```

`show global variables like "secure_file_priv";`

```
select * from words group by cc_word having count(*) >= 2;
select * from words where cc_word = 2238777324;

SELECT CONCAT(table_schema,'.',table_name) AS 'Table Name', CONCAT(ROUND(table_rows/1000000,4),'M') AS 'Number of Rows', CONCAT(ROUND(data_length/(1024*1024*1024),4),'G') AS 'Data Size', CONCAT(ROUND(index_length/(1024*1024*1024),4),'G') AS 'Index Size', CONCAT(ROUND((data_length+index_length)/(1024*1024*1024),4),'G') AS'Total'FROM information_schema.TABLES WHERE table_schema LIKE 'example';
```

```
explain select * from words where word = "preconfer";
```


## 2022年08月20日

复制表结构
`CREATE  TABLE IF NOT EXISTS tb_base_like (LIKE tb_base);`

复制表结构和数据

## 2022年08月13日15:29:14

- [x] 测试net-http-client POST的功能
	- [x] 建立http-server
	- [x] 打印请求方法, 参数
	- [x] 需要CURL命令调试

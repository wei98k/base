
- [ ] 打印文章 <MySQL的B-Tree索引底层结构以及具体实现原理详解>


## 2022年08月14日08:30:42

~整理时间: 2022年08月14日16:54:37 位置:blog~

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

## 2022年08月13日15:29:14

- [x] 测试net-http-client POST的功能
	- [x] 建立http-server
	- [x] 打印请求方法, 参数
	- [x] 需要CURL命令调试


# MySQL Crash Course

书籍《MySQL Crash Course》 译名：MySQL必知必会

[《MySQL Crash Course》](https://forta.com/books/0672327120/)


## 目录文件说明

- create.sql populate.sql 是书籍提供的示例SQL
- ch开头目录书籍示例SQL练习

## Docker操作

```
sudo docker build -t mysql56 -f Docker56 .
```

```
sudo docker run -e MYSQL_ROOT_PASSWORD=root123 -dit mysql56  bash
```
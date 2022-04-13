
# 查询单列
SELECT prod_name FROM products;

# 查询多列
select prod_id, prod_name, prod_price from products;

# 查询全部列(使用通配符)
select * from products;

# 去重复查询列
select distinct vend_id from products;

# 限制行数
select prod_name from products limit 5;

# 指定第5行开始, 限制5行
select prod_name from products limit 5,5;

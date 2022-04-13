-- 使用WHERE子句过滤数据

SELECT prod_name, prod_price
FROM products
WHERE prod_price = 2.50;

-- 检查单个值

SELECT prod_name, prod_price
FROM products
WHERE prod_name = 'fuses';

-- eg.1 列出价格小于10美元的所有产品

SELECT prod_name, prod_price
FROM products
WHERE prod_price < 10;

-- eg.2 列出小于等于10美元的所有产品

SELECT prod_name, prod_price
FROM products
WHERE prod_price <= 10;

-- 不匹配检查 列出不是由供应商1003制造的所有产品

SELECT vend_id, prod_name
FROM products
WHERE vend_id <> 1003;

SELECT vend_id, prod_name
FROM products
WHERE vend_id != 1003;

-- 检查范围 之间的数据

SELECT prod_name, prod_price
FROM products
WHERE prod_price BETWEEN 5 AND 10;

-- 空值检查

SELECT prod_name
FROM products
WHERE prod_price IS NULL;

SELECT cust_id
FROM customers
WHERE cust_email IS NULL;


























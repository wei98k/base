-- 数据过滤

-- AND操作符号

SELECT prod_id, prod_price, prod_name
FROM products
WHERE vend_id = 1003 AND prod_price <= 10;

-- OR操作符号

SELECT prod_name, prod_price
FROM products
WHERE vend_id = 1002 OR vend_id = 1003;

-- 计算次序

SELECT prod_name, prod_price
FROM products
WHERE vend_id = 1002 OR vend_id = 1003 AND prod_price >= 110;

SELECT prod_name, prod_price
FROM products
WHERE (vend_id = 1002 OR vend_id = 1003) AND prod_price >= 10;

-- IN操作符号

SELECT prod_name, prod_price
FROM products
WHERE vend_id IN(1002, 1003)
ORDER BY prod_name;

SELECT prod_name, prod_price
FROM products
WHERE vend_id = 1002 OR vend_id = 1003
ORDER BY prod_name;

-- NOT操作符

SELECT prod_name, prod_price
FROM products
WHERE vend_id NOT IN(1002,1003)
ORDER BY prod_name;



# 排序

SELECT prod_name
FROM products
ORDER BY prod_name;


# 下面的代码检索3个列，并按其中两个列对结果进行排序——首先按价格，然后再按名称排序。

SELECT prod_id, prod_price, prod_name
FROM products
ORDER BY prod_price, prod_name;

-- 指定排序方式 升序或降序

SELECT prod_id, prod_price, prod_name
FROM products
ORDER BY prod_price DESC;

-- 对多个列排序 DESC只对前面一列起作用 默认就是ASC

SELECT prod_id, prod_price, prod_name
FROM products
ORDER BY prod_price DESC, prod_name;

-- 使用ORDER BY 和 LIMIT的组合找到一个最大或最小的值

SELECT prod_price
FROM products
ORDER BY prod_price DESC
LIMIT 1;
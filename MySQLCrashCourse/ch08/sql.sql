-- LIKE操作符号

-- 百分号(%)通配符 多个匹配

SELECT prod_id, prod_name
FROM products
WHERE prod_name LIKE 'jet%';


SELECT prod_id, prod_name
FROM products
WHERE prod_name LIKE '%anvil%';


SELECT prod_name
FROM products
WHERE prod_name LIKE 's%e';

-- 下划线(_)通配符 单个匹配

SELECT prod_id, prod_name
FROM products
WHERE prod_name LIKE '_ ton anvil';

SELECT prod_id, prod_name
FROM products
WHERE prod_name LIKE '% ton anvil';























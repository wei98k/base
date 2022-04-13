SELECT Concat(vend_name, '(', vend_country, ')')
FROM vendors
ORDER BY vend_name;


SELECT cust_name, cust_contact
FROM customers
WHERE cust_contact = 'Y. Lie';


SELECT cust_name, cust_contact
FROM customers
WHERE Soundex(cust_contact) = Soundex('Y Lie');

SELECT cust_id, order_num
FROM orders
WHERE order_date = '2005-09-01';

































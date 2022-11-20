DROP TABLE IF EXISTS `merchant`;
CREATE TABLE `merchant`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `app_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `app_secret` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mch_id` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `mch_key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `user_id` int(9) NULL DEFAULT NULL,
  `status` int(3) NULL DEFAULT NULL,
  `logo_id` int(11) NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

CREATE  TABLE IF NOT EXISTS test.city (LIKE sakila.city); 

CREATE  TABLE IF NOT EXISTS test.city SELECT * FROM sakila.city;

insert into test.city select * from test.city;

update test.city set city = (select city from sakila.city order by rand() limit 1);

-- 分组统计
select count(*) as cnt, city from test.city group by city order by cnt desc limit 10;
-- 分析字符前缀
select count(*) as cnt, left(city,3) as pref from test.city group by pref order by cnt desc limit 10;
select count(*) as cnt, left(city,8) as pref from test.city group by pref order by cnt desc limit 10;

select count(distinct left(city, 3))/count(*) as sel3,
      count(distinct left(city, 4))/count(*) as sel4,
      count(distinct left(city, 5))/count(*) as sel5,
      count(distinct left(city, 6))/count(*) as sel6,
      count(distinct left(city, 7))/count(*) as sel7,
      count(distinct left(city, 8))/count(*) as sel8,
      count(distinct left(city, 9))/count(*) as sel9
from test.city;
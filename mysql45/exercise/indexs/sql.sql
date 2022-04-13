
-- 联合索引失效测试 -----------

-- 创建索引调试数据库
create database ex_index;

-- 联合索引失效测试
create table `db_index` (
    `id` int unsigned not null auto_increment,
    `a` varchar(100) not null default '',
    `b` varchar(50) not null default '',
    `c` char(10) not null default '',
    `d` int(10),
    primary key (`id`)
) engine=InnoDB;

-- 创建测试数据

DELIMITER $$
CREATE PROCEDURE insert_db_index(max_num INT)
BEGIN
DECLARE i INT DEFAULT 0;
SET autocommit = 0;
REPEAT
SET i = i + 1;
INSERT INTO db_index(a, b, c, d) VALUES(rand_string(8), rand_string(20), rand_string(8), rand_num(1, 500000));
UNTIL i = max_num
END REPEAT;
COMMIT;
END $$

DELIMITER ;
CALL insert_db_index(10000);

-- 删除
DELIMITER ;
drop PROCEDURE insert_dept;



ALTER table db_index ADD INDEX sindex(a,b,c);

show index from db_index;

-- 

explain select * from db_index where a="a" and b="b";

explain select * from db_index where c="ZwvMcvNh" and a="xuHR3whF";

show databases;
desc ex_index._index;

use ex_index;

CREATE TABLE `user_test` (
  `user_id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `login_name` VARCHAR(30) NOT NULL COMMENT '登录账号',
  `email` VARCHAR(50) DEFAULT '' COMMENT '用户邮箱',
  `phonenumber` VARCHAR(11) DEFAULT '' COMMENT '手机号码',
  `sex` CHAR(1) DEFAULT '0' COMMENT '用户性别（0男 1女 2未知）',
  `password` VARCHAR(50) DEFAULT '' COMMENT '密码',
  `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE=INNODB DEFAULT CHARSET=utf8mb4 COMMENT='测试'

use ex_index;

alter table user_test add INDEX `sindex` (`email`,`phonenumber`);

-- 769003@qq.com | 13566022421
use ex_index;
explain select * from user_test 
        where  email = '769003@qq.com'
        and  phonenumber = '13556626230';

use ex_index;
explain select * from user_test 
        where phonenumber = '13556626230'
        and email = '769003@qq.com';


show tables;

use ex_index;
DELIMITER $$
DROP PROCEDURE IF EXISTS `test_proc`$$
CREATE PROCEDURE `test_proc`()
BEGIN
	DECLARE num INT DEFAULT 10000;
	DECLARE i INT DEFAULT 1;
	WHILE i<num DO
		INSERT INTO user_test(`login_name`,`email`,`phonenumber`,`sex`,`password`) 
		VALUES(CONCAT('用户',i),CONCAT(FLOOR(RAND()*((999999-111111)+111111)),'@qq.com'),
		CONCAT('135',FLOOR(RAND()*((99999999- 11111111)+11111111))), FLOOR(RAND()*2),UUID());
		SET i = i+1;
	END WHILE;
END$$
DELIMITER ;
CALL test_proc();

use ex_index;
CREATE TABLE `test_innodb` (
    `id` INT (11) NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR (20) NOT NULL,
    `group_id` INT (11) NOT NULL,
    `create_time` datetime NOT NULL,
    PRIMARY KEY (`id`),
    KEY `index_user_id` (`user_id`) USING HASH
) ENGINE = INNODB AUTO_INCREMENT = 1 DEFAULT CHARSET = utf8;

use ex_index;

delimiter $$
CREATE FUNCTION rand_string(n int) RETURNS varchar(255) 
begin        
  declare chars_str varchar(100) 
  default "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
  declare return_str varchar(255) default "";        
  declare i int default 0;
  while i < n do        
      set return_str=concat(return_str,substring(chars_str,floor(1+rand()*62),1));
      set i= i+1;        
  end while;        
  return return_str;    
end $$


delimiter ;

DELIMITER $$
CREATE FUNCTION rand_num(from_num INT, to_num INT) RETURNS INT(11)
BEGIN
	DECLARE i INT DEFAULT 0;
	SET i = FLOOR(from_num + RAND() * (to_num - from_num +1));
	RETURN i;
END $$
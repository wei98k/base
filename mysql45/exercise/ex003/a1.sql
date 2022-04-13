
CREATE DATABASE ex003;

use ex003;

## -----------

create table T(c int) engine=InnoDB;

insert into T(c) values(1);

show variables like 'transaction_isolation';

set SESSION transaction_isolation="READ-COMMITTED"; #修改当前会话级别参数

select * from information_schema.innodb_trx where TIME_TO_SEC(timediff(now(), trx_started))>60;
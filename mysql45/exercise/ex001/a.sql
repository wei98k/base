create database db1;

use db1;

create table if not exists t (`id` int not null auto_increment, primary key (`id`)) engine=InnoDB default charset=utf8;

select SQL_CACHE * from t where id=10;
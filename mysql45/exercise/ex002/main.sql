create table T(ID int primary key, c int);

update T set c=c+1 where ID=2;

create table T(c int) engine=InnoDB;

insert into T(c) values(1);

show variables like 'transaction_isolation';


create database db1;

use db1;

create table t(id int primary key, a int, b int, index(a))engine=innodb;

delimiter ;;
	create procedure iddata()
	begin
		declare i int;
		set i=1;
		while(i<=1000)do
			insert into t values(i,i,i);
			set i=i+1;
		end while;
	end;;
delimiter ;
call iddata();

create database db2;
create table db2.t like db1.t
create table if not exists tbl_client_watchlist (
oid integer auto_increment,
user_id varchar(10),
primary key (oid)
);

drop view vw_todays_watchlist;

create view vw_todays_watchlist as 
select user_id, username, nationality, telephone, location, rate_1_hour from tbl_client 
where user_id in (select user_id from tbl_client_watchlist) and date(refresh_time) = date(now());


create table if not exists tbl_order (
	oid integer not null auto_increment,
	user_id varchar(10),
	username varchar(20),
	location varchar(200),
    region varchar(100),
	order_reference varchar(50),
	date_of_event date,
	time_of_event time,
	creation_date datetime,
	modification_date datetime,
	duration integer,
	rate integer,
	deductions integer,
	surplus integer,
	price integer, 			-- the final price
	status varchar(20),
	notes text,
	primary key (oid)
);

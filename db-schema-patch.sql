create table if not exists tbl_client_watchlist (
oid integer auto_increment,
user_id varchar(10),
primary key (oid)
);

drop view vw_todays_watchlist;

create view vw_todays_watchlist as 
select user_id, username, nationality, telephone, location, rate_1_hour from tbl_client 
where user_id in (select user_id from tbl_client_watchlist) and date(refresh_time) = date(now());
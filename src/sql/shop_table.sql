-- `shops`
create database if not exists winter_examination_database;
use winter_examination_database;
create table if not exists `shops`
(
    `shop_id` bigint not null auto_increment comment '主键' primary key,
    `shop_owner_id` bigint not null comment '店长',
    `shop_name` varchar(30) not null comment '商店名',
    `shop_is_delete` varchar(3) default '0' not null comment '是否删除'
) comment '`shops`';

insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('1', '1', '东北理工大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('2', '2', '东南农业大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('3', '3', '东北体育大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('4', '4', '东南农业大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('5', '5', '南艺术大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('6', '6', '西南农业大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('7', '7', '东北经贸大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('8', '8', '南理工大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('9', '9', '西北体育大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('10', '10', '东北大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('11', '11', '西体育大学');
insert into `shops` (`shop_id`, `shop_owner_id`, `shop_name`) values ('12', '12', '西北农业大学');
-- `shops`
create database if not exists winter_examination_database;
use winter_examination_database;
create table if not exists `shops`
(
    `shop_id` bigint not null auto_increment comment '主键' primary key,
    `shop_owner` varchar(20) not null comment '店长',
    `shop_name` varchar(30) not null comment '商店名',
    `shop_is_delete` varchar(3) default '0' not null comment '是否删除'
) comment '`shops`';

insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('1', '钟鹤轩', '东北理工大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('2', '孙天磊', '东南农业大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('3', '阎烨霖', '东北体育大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('4', '万俊驰', '东南农业大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('5', '蒋峻熙', '南艺术大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('6', '谢健柏', '西南农业大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('7', '高鸿煊', '东北经贸大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('8', '冯天磊', '南理工大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('9', '魏博文', '西北体育大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('10', '沈远航', '东北大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('11', '周俊驰', '西体育大学');
insert into `shops` (`shop_id`, `shop_owner`, `shop_name`) values ('12', '于琪', '西北农业大学');
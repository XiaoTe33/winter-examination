-- `orders`
create database if not exists winter_examination_database;
use winter_examination_database;
create table if not exists `orders`
(
    `order_id` bigint not null auto_increment comment '主键' primary key,
    `order_buyer_id` bigint not null comment '客户',
    `order_solder_id` bigint not null comment '商店',
    `order_goods_id` bigint not null comment '商品',
    `order_time` varchar(50) not null comment '下单时间',
    `order_address` varchar(100) not null comment '收货地址',
    `order_status` varchar(1) default '0' not null comment '订单状态'
) comment '`orders`';

insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (1, '1', '2', '2', '2022-04-10 20:20:43', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (2, '1', '2', '2', '2022-12-02 10:50:41', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (3, '1', '2', '2', '2022-07-01 18:03:06', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (4, '1', '2', '2', '2022-05-17 21:07:23', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (5, '1', '2', '2', '2022-12-19 06:27:29', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (6, '1', '2', '2', '2022-12-01 07:54:54', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (7, '1', '2', '2', '2022-11-12 04:20:08', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (8, '1', '3', '2', '2022-01-02 20:32:27', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (9, '1', '4', '2', '2022-03-23 16:53:21', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (10, '2', '2', '3', '2022-08-23 20:42:55', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (11, '2', '1', '3', '2022-09-11 19:14:44', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (12, '2', '2', '3', '2022-06-29 12:04:08', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (13, '2', '2', '3', '2022-07-07 02:50:30', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (14, '2', '3', '3', '2022-04-29 02:25:43', 'CQUPT');
insert into `orders` (`order_id`, `order_buyer_id`, `order_solder_id`, `order_goods_id`, `order_time`, `order_address`) values (15, '2', '1', '3', '2022-10-06 23:11:50', 'CQUPT');
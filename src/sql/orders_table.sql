-- `orders`
create database if not exists winter_examination_database;
use winter_examination_database;
create table if not exists `orders`
(
    `order_id` bigint not null auto_increment comment '主键' primary key,
    `order_buyer` varchar(20) not null comment '客户',
    `order_solder` varchar(30) not null comment '商店',
    `order_goods_id` varchar(100) not null comment '商品',
    `order_time` varchar(50) not null comment '下单时间'
) comment '`orders`';

insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (1, '程天翊', '东北农业大学', '钱越泽', '2022-04-10 20:20:43');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (2, '魏立轩', '中国农业大学', '吴浩宇', '2022-12-02 10:50:41');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (3, '郑瑞霖', '中国技术大学', '韩子涵', '2022-07-01 18:03:06');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (4, '郝远航', '东南科技大学', '袁昊天', '2022-05-17 21:07:23');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (5, '彭伟祺', '南体育大学', '罗文博', '2022-12-19 06:27:29');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (6, '邱烨伟', '中国科技大学', '于明辉', '2022-12-01 07:54:54');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (7, '潘昊然', '西北艺术大学', '朱熠彤', '2022-11-12 04:20:08');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (8, '武正豪', '北大学', '龙子涵', '2022-01-02 20:32:27');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (9, '丁峻熙', '东南艺术大学', '高绍齐', '2022-03-23 16:53:21');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (10, '洪明', '西南科技大学', '田智宸', '2022-08-23 20:42:55');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (11, '陶健柏', '中国经贸大学', '冯浩轩', '2022-09-11 19:14:44');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (12, '苏越泽', '东北理工大学', '谢天宇', '2022-06-29 12:04:08');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (13, '郭明杰', '西科技大学', '田熠彤', '2022-07-07 02:50:30');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (14, '邱子骞', '南农业大学', '杜涛', '2022-04-29 02:25:43');
insert into `orders` (`order_id`, `order_buyer`, `order_solder`, `order_goods_id`, `order_time`) values (15, '蒋明轩', '东南技术大学', '杨思淼', '2022-10-06 23:11:50');
-- `goods`
create database if not exists winter_examination_database;
use winter_examination_database;
create table if not exists `goods`
(
    `goods_id` bigint not null auto_increment comment '主键' primary key,
    `goods_is_deleted` tinyint default 0 not null comment '是否删除(0-未删, 1-已删)',
    `goods_name` varchar(256) not null comment '商品名',
    `goods_kind` varchar(256) default '其他' not null comment '商品种类',
    `goods_price` float not null comment '商品价格',
    `goods_sold_amount` bigint default 0 not null comment '销量',
    `goods_score` varchar(20) default '5.0' not null comment '商品评分',
    `goods_shop_id` varchar(20) not null comment '所属店铺'
) comment '`goods`';

insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('谭鸿煊', '服装', 67485.86, 7, '3.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('潘越泽', '电脑', 72182.305, 942, '1.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('朱智辉', '其他', 35518.52, 875, '4.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('朱明', '手机', 81894.5, 4542, '4.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('金哲瀚', '服装', 52972.453, 443856902, '5.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('何雪松', '电脑', 61282.73, 70, '4.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('宋智渊', '手机', 88197.26, 276089, '3.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('孔志泽', '服装', 16714.006, 74210900, '3.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('陶志泽', '手机', 92474.45, 3, '5.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('白俊驰', '服装', 4639.6436, 6005755659, '2.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('唐鹏煊', '手机', 24332.268, 310159, '4.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('韦峻熙', '手机', 56186.734, 146008, '3.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('何鹏', '其他', 72338.26, 275998, '4.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('张梓晨', '服装', 4382.7114, 8, '1.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('杨天宇', '电脑', 29436.725, 2, '3.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('贾伟祺', '食品', 87675.766, 1712, '4.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('唐鹤轩', '电脑', 80144.164, 41794334, '2.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('吕智渊', '其他', 79619.73, 30, '4.0');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('夏胤祥', '电脑', 50441.91, 9, '1.5');
insert into `goods` (`goods_name`, `goods_kind`, `goods_price`, `goods_sold_amount`, `goods_score`) values ('徐烨磊', '文具', 15774.721, 5, '4.5');
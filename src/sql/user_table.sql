-- `users`
create database if not exists winter_examination_database;
use winter_examination_database;
create table if not exists `users`
(
    `user_id` bigint not null auto_increment comment 'id' primary key,
    `username` varchar(20) not null comment '用户名',
    `password` varchar(256) not null comment 'password',
    `phone` varchar(11) default '' comment '手机号',
    `email` varchar(256) default '' comment '邮箱',
    `money` varchar(256) default '0.00' not null comment '余额',
    `photo` varchar(256) default '' comment '头像',
    `shopping_car` varchar(256) default '{}' comment '购物车',
    `address` varchar(256) default '' comment '收货地址'
) comment '`users`';

insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (1, '柳亦钦', '8m', '15179127087', '重庆邮电大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (2, '洪越彬', 'eo', '17089517558', '西南体育大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (3, '郑钰轩', '6q8U', '18432054977', '东南艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (4, '马鹏飞', '7RX4', '15758760380', '西北科技大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (5, '阎浩然', 'o9tX0', '15753750569', '西理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (6, '邹笑愚', 'acNRf', '17606854200', '西南大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (7, '高炫明', 'gH8H', '15247892393', '东技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (8, '杨鹤轩', 'j8Mu', '15760732736', '东技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (9, '覃擎苍', 'jry3X', '15059117940', '西北理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (10, '梁炎彬', 'ThOJY', '17056566511', '南大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (11, '丁熠彤', 'Qkul', '17757829514', '东南大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (12, '蔡聪健', 'N3', '15086133867', '东北体育大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (13, '黄伟宸', 'oO', '17227563378', '东农业大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (14, '毛乐驹', 'Ij', '17629180448', '西北经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (15, '谢思远', 'qWi', '15501703914', '西北艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (16, '丁琪', 'qQa', '18922923542', '东北科技大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (17, '方越彬', 'vzS', '17676998197', '西南技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (18, '贾熠彤', 'bOLlC', '15379873939', '北大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (19, '朱烨伟', 'D6LRg', '17601547734', '南理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (20, '彭子涵', 'VW7', '15149459035', '东理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (21, '汪文博', '7s', '17566495083', '东经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (22, '钟思源', 'qD', '17327650744', '中国技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (23, '李雪松', 'ySI', '15019461574', '西南技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (24, '梁明轩', 'qPt9', '15046956678', '中国技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (25, '段晓啸', 'x9bUa', '15555693473', '西理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (26, '曹智渊', 'je', '15807983225', '西南技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (27, '马子骞', 'RXb', '17759630362', '东南大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (28, '吴思', '7rvZO', '17272477832', '西南艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (29, '钱熠彤', 'hg', '15554975076', '西南经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (30, '武明轩', 'wvCa1', '15544383511', '东南理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (31, '孔雨泽', 'lwE', '15896519747', '西北艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (32, '黎伟祺', 'Om', '15155217857', '北艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (33, '秦天翊', 'hywLR', '15796810279', '东科技大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (34, '郑笑愚', 'VSGQ', '14759470588', '西北技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (35, '杜鹏涛', 'UCyyj', '14761258461', '东北艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (36, '郝钰轩', 'jnkh', '15099389295', '东艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (37, '钱子涵', 'Vk', '17734229704', '东南艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (38, '龙锦程', 'cjHL', '15928972488', '东南科技大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (39, '曹嘉懿', 'gi9Gc', '15368968279', '南大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (40, '尹子涵', '8Q9Y', '17124035019', '南经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (41, '赵君浩', 'Px', '15690695753', '北理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (42, '邱昊然', 'tW', '17320761048', '西北技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (43, '赵文轩', 'Lj8MP', '15614555619', '中国理工大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (44, '韩明', 'NAr', '17681276187', '东北技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (45, '严智渊', 'FxmU', '15828269694', '东南经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (46, '谢果', 'Cx', '15218683461', '东经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (47, '罗鸿煊', 'FZhdZ', '18398076901', '东北体育大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (48, '卢文昊', '5jivm', '15076839632', '中国农业大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (49, '贾天翊', 'sYz5', '15900106066', '西技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (50, '郭雪松', 'zkqB', '15621022095', '西技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (51, '何旭尧', '3q', '17051011783', '西艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (52, '唐思淼', 'IzGJ', '15180378012', '东北大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (53, '段烨伟', 'Fjmb', '15766136291', '东经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (54, '廖博超', 'N8', '14775062999', '西北经贸大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (55, '谭文博', 'YbeE', '17702495550', '西艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (56, '卢伟诚', 'Gz', '15068723824', '西北技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (57, '方修洁', 'sJe', '15593361491', '东技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (58, '唐昊强', 'we', '17022242248', '东科技大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (59, '黄航', 'L4Ty', '15699533409', '中国艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (60, '彭明', 'Lp8V', '17109938354', '北农业大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (61, '范修杰', 'W0zPx', '13091748879', '东北农业大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (62, '汪靖琪', 'kynX', '17385808417', '西艺术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (63, '姚绍辉', 'JTYK', '17884037476', '南体育大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (64, '魏睿渊', 'po', '15345038472', '南技术大学');
insert into `users` (`user_id`, `username`, `password`, `phone`, `address`) values (65, '谭思远', 'PbE', '15944973170', '中国科技大学');


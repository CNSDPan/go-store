CREATE TABLE `users` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `user_id` bigint(20) NOT NULL COMMENT '用户IID',
 `token` varchar(18)  DEFAULT '' COMMENT 'token',
 `status` tinyint(1) DEFAULT '1' COMMENT '1=启用 2=禁用',
 `name` varchar(50)  DEFAULT '' COMMENT '昵称',
 `fund` bigint(20) unsigned DEFAULT '0' COMMENT '用户资金,入库*1000【1000 = 1元】',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `unx_user_token` (`user_id`,`token`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
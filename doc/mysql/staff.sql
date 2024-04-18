CREATE TABLE `staff` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `staff_sn` bigint(20) NOT NULL COMMENT '用户IID',
 `user` varchar(255)  NOT NULL COMMENT '账号',
 `password` varchar(255)  NOT NULL,
 `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1=启用 0=禁用',
 `name` varchar(50)  NOT NULL DEFAULT '' COMMENT '昵称',
 `cn_name` varchar(50)  NOT NULL DEFAULT '' COMMENT '真实姓名',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `unx_user_phone` (`staff_sn`,`user`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
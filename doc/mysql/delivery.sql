CREATE TABLE `delivery` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`delivery_sn` bigint(20) NOT NULL COMMENT '配送物流编号IID',
`name` varchar(191) DEFAULT '' COMMENT '配送公司',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `delivery_sn` (`delivery_sn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='配送物流';
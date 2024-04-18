CREATE TABLE `store` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`store_sn` bigint(20) NOT NULL COMMENT '店铺IID',
`name` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '店铺',
`code` varchar(191) DEFAULT '' COMMENT '店铺编号',
`remark` varchar(255) DEFAULT '' COMMENT '备注',
`address_country` int(10) DEFAULT '0' COMMENT '国家',
`address_province` int(10) DEFAULT '0' COMMENT '省',
`address_city` int(10) DEFAULT '0' COMMENT '市',
`address_district` int(10) DEFAULT '0' COMMENT '区',
`address_detail` text COLLATE utf8mb4_unicode_ci COMMENT '快递详细地址',
`address_post_code` varchar(191) COLLATE utf8mb4_unicode_ci DEFAULT '' COMMENT '邮编',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `unx_store_code` (`store_sn`,`code`) USING BTREE,
KEY `idx_add` (`address_country`,`address_province`,`address_city`,`address_district`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='店铺';
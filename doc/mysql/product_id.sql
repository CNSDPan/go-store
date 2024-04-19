CREATE TABLE `product` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`product_id` bigint(20) NOT NULL COMMENT '商品IID',
`title` varchar(255) DEFAULT '' COMMENT '商品中文名称',
`image` varchar(500) DEFAULT '' COMMENT '商品图片',
`status` tinyint(1) unsigned DEFAULT '1' COMMENT '状态:1-正在销售、2-新品、2-爆款、9-停止销售',
`price` bigint(20) unsigned DEFAULT '0' COMMENT '价钱',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `unx_product` (`product_id`) USING BTREE,
KEY `idx_status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='sku商品表';
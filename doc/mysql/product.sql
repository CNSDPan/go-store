CREATE TABLE `product` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`product_sn` bigint(20) NOT NULL COMMENT '产品IID',
`title` varchar(500) NOT NULL COMMENT '商品中文名称',
`spu` varchar(50) NOT NULL DEFAULT '' COMMENT '商品货号（SPU）',
`image` varchar(500) NULL DEFAULT '' COMMENT '商品图片',
`status` tinyint(1) unsigned DEFAULT '1' COMMENT '状态:1-正常，9-删除',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `unx_product` (`product_sn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='spu商品表';
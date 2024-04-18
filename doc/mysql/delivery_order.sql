CREATE TABLE `delivery_order` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `delivery_order_sn` bigint(20) NOT NULL COMMENT '配送编号',
 `order_sn` bigint(20) DEFAULT NULL COMMENT '订单编号IID',
 `detail_sn` bigint(20) DEFAULT NULL COMMENT '明细编号IID',
 `cn_name` varchar(191) DEFAULT '' COMMENT '骑手',
 `phone` varchar(191) DEFAULT '' COMMENT '骑手电话',
 `inspected` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '配送数量',
 `type` tinyint(1) unsigned DEFAULT '1' COMMENT '配送类型：1-首次配送、2-追加配送',
 `delivery_sn` tinyint(1) unsigned DEFAULT '1' COMMENT '配送渠道',
 `express_fee` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '配送费，入库*1000【1000 = 1元】',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `delivery_sn` (`delivery_order_sn`) USING BTREE,
 KEY `idx_order_detail` (`order_sn`,`detail_sn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='订单配送明细';
CREATE TABLE `pay_order` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `pay_id` bigint(20) NOT NULL COMMENT '支付单编号IID',
 `order_id` bigint(20) DEFAULT '0' COMMENT '订单编号IID',
 `type` tinyint(1) unsigned DEFAULT '0' COMMENT '支付来源：1-微信、2-支付宝、3-美团、4-第三方',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `pay_id` (`pay_id`) USING BTREE,
 KEY `idx_order_type` (`order_id`,`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='支付订单列表';
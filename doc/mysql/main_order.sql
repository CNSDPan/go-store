CREATE TABLE `main_order` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`order_sn` bigint(20) NOT NULL COMMENT '订单编号',
`store_id` bigint(20) DEFAULT NULL COMMENT '店铺IID',
`status` tinyint(2) unsigned DEFAULT '1' COMMENT '订单状态：1-待支付、2-待接单、3-待配送、4-待送达、11-取消订单、12-失效订单、20-已完成',
`pay_status`tinyint(2) unsigned DEFAULT '1' COMMENT '支付状态:1-待支付、2-取消支付、3-支付超时、11-支付失败、20-已支付',
`pay_time` bigint(20) DEFAULT '0' COMMENT '支付时间,毫秒',
`pay_timeout` bigint(20) DEFAULT '0' COMMENT '支付有效时间,毫秒',
`pay_time_close` bigint(20) DEFAULT '0' COMMENT '支付“取消|失效 ”时间,毫秒',
`total` bigint(20) unsigned DEFAULT '0' COMMENT '订单总价,入库*1000【1000 = 1元】',
`quantity` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '商品总数量',
`inspected` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '出单总数量',
`store_receive_status` tinyint(1) unsigned DEFAULT '0' COMMENT '商家接单状态：1-待接单、2-待出单、20-已出单',
`store_receive_time` bigint(20) DEFAULT '0' COMMENT '接单时间,毫秒',
`store_issue_time` bigint(20) DEFAULT '0' COMMENT '出单时间,毫秒',
`ride_receive_status` tinyint(1) unsigned DEFAULT '0' COMMENT '骑手接单状态：1-待接单、2-已接单、11-取消送单、20-已送达',
`ride_store_receive_time` bigint(20) DEFAULT '0' COMMENT '骑手接单时间,毫秒',
`ride_store_issue_time` bigint(20) DEFAULT '0' COMMENT '骑手送达时间,毫秒',
`operate_time` bigint(20) DEFAULT '0' COMMENT '订单操作时间,毫秒',
`finish_time` bigint(20) DEFAULT '0' COMMENT '订单完成时间,毫秒',
`remark` varchar(255) DEFAULT '' COMMENT '订单备注',
`staff_sn` bigint(20) DEFAULT '0' COMMENT '负责人IID',
`address_name` varchar(191) DEFAULT '' COMMENT '收货人姓名',
`address_phone` varchar(191) DEFAULT '' COMMENT '收货人电话',
`address_country` int(10) DEFAULT '0' COMMENT '国家',
`address_province` int(10) DEFAULT '0' COMMENT '省',
`address_city` int(10) DEFAULT '0' COMMENT '市',
`address_district` int(10) DEFAULT '0' COMMENT '区',
`address_detail` text COMMENT '详细地址',
`address_post_code` varchar(191) DEFAULT '' COMMENT '邮编',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `order_sn` (`order_sn`) USING BTREE,
KEY `idx_store` (`store_id`) USING BTREE,
KEY `idx_status` (`status`) USING BTREE,
KEY `idx_staff` (`staff_sn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='订单列表';
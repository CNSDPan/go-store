CREATE TABLE `purchase_order_detail` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `detail_sn` bigint(20) NOT NULL COMMENT '明细编号IID',
 `order_sn` bigint(20) DEFAULT NULL COMMENT '订单编号IID',
 `product_variant_sn` mediumint(8) unsigned NOT NULL DEFAULT '0' COMMENT '商品IID：product_variant.variant_sn',
 `total` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品总价入库*1000【1000 = 1元】',
 `quantity` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
 `inspected` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '出单数量',
 `price` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '商品单价入库*1000【1000 = 1元】',
 `return_quantity` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '退货数量',
 `repair_inspected` int(8) unsigned NOT NULL DEFAULT '0' COMMENT '补出单数量',
 `remark` varchar(255) DEFAULT '' COMMENT '备注',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `unx_detail_order` (`detail_sn`,`order_sn`) USING BTREE,
 KEY `idx_variant_sn` (`product_variant_sn`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='订单商品明细表';
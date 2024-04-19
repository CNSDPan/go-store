CREATE TABLE `fund_expenses` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`expenses_id` bigint(20) NOT NULL COMMENT '支出IID',
`user_id` bigint(20) DEFAULT '0' COMMENT '用户IID',
`type` tinyint(1) DEFAULT '0' COMMENT '支出类型：1-订单、2-红包',
`before` bigint(20) unsigned DEFAULT '0' COMMENT '扣除前',
`after` bigint(20) unsigned DEFAULT '0' COMMENT '扣除后',
`created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
`updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
PRIMARY KEY (`id`) USING BTREE,
UNIQUE KEY `unx_expenses` (`expenses_id`) USING BTREE,
KEY `idx_user` (`user_id`) USING BTREE,
KEY `idx_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='资金支出流水';

CREATE TABLE `fund_income` (
 `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
 `income_id` bigint(20) NOT NULL COMMENT '收入IID',
 `user_id` bigint(20) DEFAULT '0' COMMENT '用户IID',
 `type` tinyint(1) DEFAULT '0' COMMENT '收入类型：2-红包',
 `before` bigint(20) unsigned DEFAULT '0' COMMENT '收入前',
 `after` bigint(20) unsigned DEFAULT '0' COMMENT '收入后',
 `created_at` datetime(3) DEFAULT NULL COMMENT '创建时间',
 `updated_at` datetime(3) DEFAULT NULL COMMENT '更新时间',
 PRIMARY KEY (`id`) USING BTREE,
 UNIQUE KEY `unx_income` (`income_id`) USING BTREE,
 KEY `idx_user` (`user_id`) USING BTREE,
 KEY `idx_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='资金收入流水';
CREATE TABLE `china_area` (
`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`code` bigint unsigned NOT NULL COMMENT '区划代码',
`name` varchar(128) NOT NULL DEFAULT '' COMMENT '名称',
`level` tinyint(1) NOT NULL COMMENT '级别1-5,省市县镇村',
`pcode` bigint DEFAULT NULL COMMENT '父级区划代码',
`category` int DEFAULT NULL COMMENT '城乡分类',
PRIMARY KEY (`id`),
UNIQUE KEY `code` (`code`) USING BTREE,
KEY `name` (`name`),
KEY `level` (`level`),
KEY `pcode` (`pcode`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='2024-行政区域';
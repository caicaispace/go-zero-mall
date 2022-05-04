CREATE TABLE `user_cart` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` INT(10) UNSIGNED NOT NULL DEFAULT '0',
	`source_id` INT(10) UNSIGNED NOT NULL DEFAULT '0',
	`source_type` INT(10) UNSIGNED NOT NULL DEFAULT '0',
	`license_type` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '授权类型(100:个人,200:企业,210企业plus,300:单张)',
	`video_rate` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '等级\\r\\n0 : 其他\\r\\n1 : 1280x720\\r\\n2 : 1920x1080\\r\\n3 : 4096x2169(4k)\\r\\n4 : 8192x4320(8k)\\r\\n5 : 2048x1080(2k)\\r\\n6 : 3840x2160(UHD 4K)\\r\\n7 : 7680x4320(UHD 8K)',
	`source_num` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '1' COMMENT '素材数量(用于购买多个)',
	`is_del` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0',
	`deleted_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`deleted_at` DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
	`created_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `user_id_source_id_source_type_license_type` (`user_id`, `source_id`, `source_type`, `license_type`) USING BTREE,
	INDEX `user_id_is_del` (`user_id`, `is_del`) USING BTREE
)
COMMENT='购物车'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;

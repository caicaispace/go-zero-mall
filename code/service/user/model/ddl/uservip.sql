CREATE TABLE `user_vip` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`user_id` INT(10) UNSIGNED NOT NULL DEFAULT '0',
	`vip_id` INT(10) UNSIGNED NOT NULL DEFAULT '0',
	`vip_type` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '对应 vip 表 type',
	`video_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '视频 id（单张购买使用）',
	`order_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '订单 id',
	`license_id` INT(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'cd_user_vip_license表id',
	`start_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '起效时间',
	`start_at` DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
	`end_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '失效时间',
	`end_at` DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
	`day_limit` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '日下载限制',
	`total_limit` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '总下载限制',
	`last_admin_user` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '最后一个操作vip的人' COLLATE 'utf8_general_ci',
	`remark` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '备注{为什么添加}' COLLATE 'utf8_general_ci',
	`is_del` TINYINT(1) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除',
	`created_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`deleted_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`deleted_at` DATETIME NOT NULL DEFAULT '0000-00-00 00:00:00',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `user_id` (`user_id`) USING BTREE,
	INDEX `user_id_vip_type_end_time` (`user_id`, `vip_type`, `end_time`) USING BTREE,
	INDEX `user_id_vip_type_video_id_end_time` (`user_id`, `vip_type`, `video_id`, `end_time`) USING BTREE,
	INDEX `order_id` (`order_id`) USING BTREE,
	INDEX `license_id` (`license_id`) USING BTREE,
	INDEX `created_time` (`created_time`) USING BTREE
)
COMMENT='用户 vip'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;

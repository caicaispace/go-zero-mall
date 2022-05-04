CREATE TABLE `vip` (
	`id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	`title` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '标题' COLLATE 'utf8_general_ci',
	`title_simple` VARCHAR(20) NOT NULL DEFAULT '' COMMENT '简写标题' COLLATE 'utf8_general_ci',
	`desc` VARCHAR(300) NOT NULL DEFAULT '' COMMENT '描述' COLLATE 'utf8_general_ci',
	`appliance` VARCHAR(300) NOT NULL DEFAULT '' COMMENT '适用范围' COLLATE 'utf8_general_ci',
	`type` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '1:免费 100:个人 200:企业 300:单张个人 400:单张企业',
	`class_type` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类 vip 类型',
	`auth_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '授权类型',
	`level` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '等级',
	`sort` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序',
	`price` DECIMAL(10,2) UNSIGNED NOT NULL DEFAULT '0.00' COMMENT '价格',
	`original_price` DECIMAL(10,2) UNSIGNED NOT NULL DEFAULT '0.00' COMMENT '原价',
	`day_limit` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '日下载限制',
	`total_limit` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '总下载限制',
	`reset_cycle` SMALLINT(5) UNSIGNED NOT NULL DEFAULT '0' COMMENT '重置周期（单位/天）',
	`parent_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父级id',
	`single_id` INT(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT 'VIP对应单张 id (0代表单个购买价格，大于0代表VIP价格(其中1代表无对应售卖的单款价格)）',
	`is_usable` TINYINT(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否可用',
	`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	`created_time` BIGINT(20) UNSIGNED NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	INDEX `class_type` (`class_type`) USING BTREE,
	INDEX `id_sort_single_id_is_usable` (`id`, `sort`, `single_id`, `is_usable`) USING BTREE,
	INDEX `level` (`level`) USING BTREE
)
COMMENT='vip 表'
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;

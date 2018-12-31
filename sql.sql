CREATE TABLE IF NOT EXISTS `user_info`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `user_name` varchar(256) NOT NULL DEFAULT '' COMMENT '用户名',
  `depart_name` varchar(256) NOT NULL DEFAULT '' COMMENT '部门名称',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户信息表'

CREATE TABLE IF NOT EXISTS `user_detail`(
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `intro` text DEFAULT NULL COMMENT '具体详细信息',
  `profile` text DEFAULT NULL COMMENT '头像',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET=utf8mb4 COMMENT '用户详细信息表'


/**
 * 表fc_order_withdraw_payment_out新增字段
 */
ALTER TABLE fc_order_withdraw_payment_out 
ADD COLUMN withdraw_status tinyint(1) NOT NULL DEFAULT 1 COMMENT '打款状态(0 未打款; 1 已打款)';

/**
 * 表daily_task新增字段
 */
ALTER TABLE daily_task 
ADD COLUMN include_game_codes text NOT NULL COMMENT '包含游戏',
ADD COLUMN exclude_game_codes text NOT NULL COMMENT '屏蔽游戏',
ADD COLUMN cycle tinyint(1) NOT NULL DEFAULT 0 COMMENT '周期(0 长期; 1 自然日; 2 自然周; 3 自然月)';

/**
 * 创建mail_template新表
 */
CREATE TABLE `mail_template` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `content` text COLLATE utf8mb4_general_ci COMMENT '模板内容',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='邮件模板表';

/**
 * 创建mail新表
 */
CREATE TABLE `mail` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `user_ids` text COLLATE utf8mb4_general_ci COMMENT '用户组',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型(0 人工发送; 1 首次登录; 2 充值成功; 3 提款失败; 4 提款成功)',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `content` text COLLATE utf8mb4_general_ci COMMENT '模板内容',
  `is_popup` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否弹窗(0 否; 1 是)',
  `is_keep` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否保留(0 否; 1 是)',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态(0 关闭; 1 开启)',
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='后台邮件表';

/**
 * 创建fc_user_mail新表
 */
CREATE TABLE `fc_user_mail` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `msg_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '用户ID',
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型(0 人工发送; 1 首次登录; 2 充值成功; 3 提款失败; 4 提款成功)',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标题',
  `content` text COLLATE utf8mb4_general_ci COMMENT '模板内容',
  `is_popup` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否弹窗(0 否; 1 是)',
  `is_keep` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否保留(0 否; 1 是)',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `read_status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '已读状态(1 未读; 2 已读)',
  `del_status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '删除状态(0 未删除; 1 已删除)',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户邮件表';

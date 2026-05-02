/**
 * 表fc_promotion_info新增字段
 */
ALTER TABLE fc_promotion_info 
ADD COLUMN h5_content text NOT NULL COMMENT '活动h5详情',
ADD COLUMN stage_content text NOT NULL COMMENT '阶段详情',
ADD COLUMN gift_style tinyint(1) NOT NULL DEFAULT 0 COMMENT '礼包样式',
ADD COLUMN recharge_balance_ratio decimal(12,2) NOT NULL DEFAULT 0 COMMENT '充值余额比',
ADD COLUMN balance decimal(12,2) NOT NULL DEFAULT 0 COMMENT '余额',
ADD COLUMN first_recharge_amount decimal(12,2) NOT NULL DEFAULT 0 COMMENT '首充金额',
ADD COLUMN bonus_amount decimal(12,2) NOT NULL DEFAULT 0 COMMENT '奖励金额',
ADD COLUMN reg_start_time datetime DEFAULT NULL COMMENT '注册开始时间',
ADD COLUMN reg_end_time datetime DEFAULT NULL COMMENT '注册结束时间',
ADD COLUMN date_range_data mediumtext NOT NULL COMMENT '时间段数据',
ADD COLUMN cycle tinyint(1) NOT NULL DEFAULT 0 COMMENT '周期(0 长期; 1 自然日; 2 自然周; 3 自然月)';

/**
 * 创建fc_user_data新表
 */
CREATE TABLE `fc_user_data` (
  `user_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '0' COMMENT '用户Id',
  `activity_data` mediumtext COLLATE utf8mb4_general_ci COMMENT '用户活动数据',
  PRIMARY KEY (`user_id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户数据表';

/**
 * 表ads_carousel新增字段
 */
ALTER TABLE ads_carousel 
ADD COLUMN jumpto VARCHAR(256) NOT NULL DEFAULT '' COMMENT '活动跳转';
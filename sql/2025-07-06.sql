/**
 * fc_bet_record数据表增加is_settled字段
 */
ALTER TABLE fc_bet_record
ADD COLUMN is_settled int(11) NOT NULL DEFAULT 0 COMMENT '是否结算';

CREATE TABLE `fc_bet_record_unsettled`  (
  `id` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL,
  `user_id` varchar(30) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '0' COMMENT '用户ID',
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户账号',
  `account` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '游戏账号',
  `player_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '三方游戏code',
  `order_sn` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '三方订单号',
  `game_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '三方游戏code',
  `bet_amount` decimal(14, 2) NOT NULL,
  `net_amount` decimal(14, 2) NOT NULL,
  `valid_betamount` decimal(14, 2) NOT NULL,
  `bet_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '投注时间',
  `settlement_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '结算时间',
  `third_bettime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '三方投注时间',
  `third_settlementtime` datetime NULL DEFAULT CURRENT_TIMESTAMP COMMENT '三方结算时间',
  `game_type` varchar(100) CHARACTER SET utf8mb3 COLLATE utf8mb3_bin NOT NULL DEFAULT '' COMMENT '游戏分类标识 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼',
  `game_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '游戏名称',
  `table_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '牌桌号',
  `venue_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '场馆code',
  `odds` decimal(12, 4) NOT NULL COMMENT '赔率',
  `odds_type` tinyint(1) NOT NULL DEFAULT 0 COMMENT '赔率类型 ',
  `version` bigint NOT NULL DEFAULT 0 COMMENT '版本',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '修改人',
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `merchant_net_amount` decimal(14, 2) NOT NULL COMMENT '商家输赢',
  `after_balance` decimal(14, 2) NOT NULL COMMENT '游戏后金额',
  `is_settled` int NOT NULL DEFAULT 0 COMMENT '是否结算',
  INDEX `user_id`(`user_id` ASC) USING BTREE,
  INDEX `order_sn`(`order_sn` ASC) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_bin COMMENT = '注单表' ROW_FORMAT = DYNAMIC;
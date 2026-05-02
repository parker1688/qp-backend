/**
 * 创建ads新表
 */
CREATE TABLE `ads_carousel` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `key` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '广告栏key',
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '广告栏名称',
  `is_carousel` tinyint(1) NOT NULL COMMENT '是否轮播(0 否; 1 是)',
  `sources` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '轮播图集数据',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态(0 关闭; 1 开启)',
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='广告栏表';
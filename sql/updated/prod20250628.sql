/**
 * 表fc_agent_domain新增字段type、customer_link、ios_link、android_link、banner_img、logo_img
 */
ALTER TABLE fc_agent_domain 
ADD COLUMN `type` TINYINT(1) NOT NULL DEFAULT 2 COMMENT '类型',
ADD COLUMN customer_link VARCHAR (255) NOT NULL DEFAULT '' COMMENT '客服链接',
ADD COLUMN ios_link VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'IOS链接',
ADD COLUMN android_link VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'Android链接',
ADD COLUMN banner_img VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'banner图',
ADD COLUMN logo_img VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'logo图';

/**
 * 表fc_merchant_link新增字段logo_img
 */
ALTER TABLE fc_merchant_link 
ADD COLUMN logo_img VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'logo图',
ADD COLUMN banner_img VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'banner图';

/**
 * 创建blacklist新表
 */
CREATE TABLE `blacklist` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `type` tinyint(1) NOT NULL COMMENT '类型(1 ip地址; 2 设备码)',
  `value` varchar(512) COLLATE utf8mb4_general_ci NOT NULL COMMENT '数据',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='黑名单表';

/**
 * 表fc_transcation新增字段
 */
ALTER TABLE fc_transcation 
ADD COLUMN funding_subtype VARCHAR (50) NOT NULL DEFAULT '' COMMENT '资金子类型';

/**
 * fc_venue_game增加gtype字段
 */
ALTER TABLE fc_venue_game 
ADD COLUMN `gtype` int DEFAULT NULL COMMENT '游戏种类';

/**
 * 创建user_group新表
 */
CREATE TABLE `user_group` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `group_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '组名',
  `data` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '组数据',
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='用户组表';

/**
 * 创建guides新表
 */
CREATE TABLE `guides` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `key` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '入口key',
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '入口名称',
  `data` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '导航数据',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='引导表';

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

/**
 * 表fc_agent_domain新增字段ios_link2、android_link2
 */
ALTER TABLE fc_agent_domain 
ADD COLUMN ios_link2 VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'IOS备用链接',
ADD COLUMN android_link2 VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'Android备用链接';

/**
 * 表fc_user_material新增字段
 */
ALTER TABLE fc_user_material 
ADD COLUMN dailybonus_data VARCHAR (512) NOT NULL DEFAULT '' COMMENT '签到数据';

/**
 * 创建daily_bonus新表
 */
CREATE TABLE `daily_bonus` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `bonus` text COLLATE utf8mb4_general_ci NOT NULL COMMENT '奖励',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='每日签到表';

/**
 * 创建splash_screen新表
 */
CREATE TABLE `splash_screen` (
  `id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `merchant_code` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '商户code',
  `logo_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'logo图片',
  `banner_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'banner图片',
  `screen_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '开屏图',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '创建人',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '修改人',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='开屏画面表';

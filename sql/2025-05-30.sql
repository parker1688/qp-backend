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
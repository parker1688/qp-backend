/**
 * 表fc_agent_domain新增字段ios_link2、android_link2
 */
ALTER TABLE fc_agent_domain 
ADD COLUMN ios_link2 VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'IOS备用链接',
ADD COLUMN android_link2 VARCHAR (255) NOT NULL DEFAULT '' COMMENT 'Android备用链接';
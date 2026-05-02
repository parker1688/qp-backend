/**
 * 表fc_transcation新增字段
 */
ALTER TABLE fc_transcation 
ADD COLUMN funding_subtype VARCHAR (50) NOT NULL DEFAULT '' COMMENT '资金子类型';
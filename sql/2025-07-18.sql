/**
 * 表fc_pay_channel新增字段
 */
ALTER TABLE fc_pay_channel 
ADD COLUMN strategy varchar(512) NOT NULL DEFAULT '[]' COMMENT '支付限单';
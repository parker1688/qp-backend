/**
 * fc_transcation数据表增加manual_related_id字段
 */
ALTER TABLE fc_transcation 
ADD COLUMN manual_related_id VARCHAR (50) NOT NULL DEFAULT '' COMMENT '人工关联id';
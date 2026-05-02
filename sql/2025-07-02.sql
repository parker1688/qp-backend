/**
 * 表ads_carousel新增字段
 */
ALTER TABLE ads_carousel 
ADD COLUMN sort int NOT NULL DEFAULT 0 COMMENT '排序';
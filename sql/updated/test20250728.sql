/**
 * 表fc_order_withdraw_payment_out新增字段
 */
ALTER TABLE fc_order_withdraw_payment_out 
ADD COLUMN withdraw_status tinyint(1) NOT NULL DEFAULT 1 COMMENT '打款状态(0 未打款; 1 已打款)';

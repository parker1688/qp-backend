ALTER TABLE `admin_user`
ADD COLUMN `total_amount` float(11) NULL COMMENT '总额度',
ADD COLUMN `limit_pertime_amount` float(11) NULL COMMENT '单次限额',
ADD COLUMN `cur_amount` float NULL COMMENT '当前使用额度';

ALTER TABLE `daily_task`
ADD COLUMN `groupid` varbinary(50) NOT NULL COMMENT '分组id' AFTER `subtype`;

CREATE TABLE `op_record`  (
    `id` varchar(50) NOT NULL,
    `user_name` varchar(50) NOT NULL,
    `user_id` varchar(50) NOT NULL,
    `merchant_code` varchar(50) NOT NULL,
    `ip` varchar(50) NOT NULL,
    `menu1` varchar(50) NOT NULL,
    `menu2` varchar(50) NOT NULL,
    `op` varchar(50) NOT NULL,
    `result` varchar(50) NOT NULL,
    `create_time` varbinary(50) NOT NULL,
    PRIMARY KEY (`id` DESC)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='后台操作日志表';
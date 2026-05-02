/**
 * 注单幂等兜底：防止同一场馆订单并发重复入库
 */
SET @has_bet_record := (
	SELECT COUNT(1)
	FROM information_schema.tables
	WHERE table_schema = DATABASE() AND table_name = 'fc_bet_record'
);
SET @has_bet_record_idx := (
	SELECT COUNT(1)
	FROM information_schema.statistics
	WHERE table_schema = DATABASE()
		AND table_name = 'fc_bet_record'
		AND index_name = 'uk_fc_bet_record_venue_order_sn'
);
SET @sql_bet_record := IF(
	@has_bet_record = 1 AND @has_bet_record_idx = 0,
	'ALTER TABLE fc_bet_record ADD UNIQUE INDEX uk_fc_bet_record_venue_order_sn (venue_code(50), order_sn(100))',
	'SELECT 1'
);
PREPARE stmt_bet_record FROM @sql_bet_record;
EXECUTE stmt_bet_record;
DEALLOCATE PREPARE stmt_bet_record;

SET @has_unsettled := (
	SELECT COUNT(1)
	FROM information_schema.tables
	WHERE table_schema = DATABASE() AND table_name = 'fc_bet_record_unsettled'
);
SET @has_unsettled_idx := (
	SELECT COUNT(1)
	FROM information_schema.statistics
	WHERE table_schema = DATABASE()
		AND table_name = 'fc_bet_record_unsettled'
		AND index_name = 'uk_fc_bet_record_unsettled_venue_order_sn'
);
SET @sql_unsettled := IF(
	@has_unsettled = 1 AND @has_unsettled_idx = 0,
	'ALTER TABLE fc_bet_record_unsettled ADD UNIQUE INDEX uk_fc_bet_record_unsettled_venue_order_sn (venue_code(50), order_sn(100))',
	'SELECT 1'
);
PREPARE stmt_unsettled FROM @sql_unsettled;
EXECUTE stmt_unsettled;
DEALLOCATE PREPARE stmt_unsettled;

ALTER TABLE menus 
ADD COLUMN open_cache tinyint NOT NULL COMMENT '1使用0不适用' AFTER show_status;
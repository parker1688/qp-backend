-- Local menu seed for super-admin
-- Generated from qp-admin-template/src/views/**/index.vue
CREATE TABLE IF NOT EXISTS `menu` (
  `id` varchar(64) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `menu_name` varchar(255) NOT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `sort` bigint DEFAULT 0,
  `role_flag` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `parent_id` varchar(64) DEFAULT NULL,
  `type` bigint DEFAULT 1,
  `api_regular` varchar(255) DEFAULT NULL,
  `perms` varchar(255) DEFAULT NULL,
  `show_status` bigint DEFAULT 1,
  `open_cache` bigint DEFAULT 0,
  `create_time` datetime DEFAULT CURRENT_TIMESTAMP,
  `create_by` varchar(64) DEFAULT NULL,
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_by` varchar(64) DEFAULT NULL,
  `locales` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DELETE FROM `menu`;
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1000','activity','activity','',10000,'activity','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1001','cfgLst','cfgLst','',9999,'cfgLst','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1002','domain','domain','',9998,'domain','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1003','finance','finance','',9997,'finance','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1004','gameManage','gameManage','',9996,'gameManage','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1005','home','home','',9995,'home','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1006','log','log','',9994,'log','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1007','merchantChannel','merchantChannel','',9993,'merchantChannel','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1008','payment','payment','',9992,'payment','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1009','permission','permission','',9991,'permission','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1010','proxy','proxy','',9990,'proxy','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1011','statistics','statistics','',9989,'statistics','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1012','systemManagement','systemManagement','',9988,'systemManagement','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1013','userLabels','userLabels','',9987,'userLabels','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1014','userManage','userManage','',9986,'userManage','',NULL,1,'(*)','','1',0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1015','benefit','benefit','',9985,'activity/benefit/index','activity/benefit/index','1000',2,'(*)','activity/benefit/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1016','checkIn','checkIn','',9984,'activity/checkIn/index','activity/checkIn/index','1000',2,'(*)','activity/checkIn/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1017','newActivity','newActivity','',9983,'activity/newActivity/index','activity/newActivity/index','1000',2,'(*)','activity/newActivity/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1018','notice','notice','',9982,'activity/notice/index','activity/notice/index','1000',2,'(*)','activity/notice/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1019','shortMessageManagement','shortMessageManagement','',9981,'activity/shortMessageManagement/index','activity/shortMessageManagement/index','1000',2,'(*)','activity/shortMessageManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1020','task','task','',9980,'activity/task/index','activity/task/index','1000',2,'(*)','activity/task/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1021','ticket','ticket','',9979,'activity/ticket/index','activity/ticket/index','1000',2,'(*)','activity/ticket/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1022','AdBannerConfiguration','AdBannerConfiguration','',9978,'cfgLst/AdBannerConfiguration/index','cfgLst/AdBannerConfiguration/index','1001',2,'(*)','cfgLst/AdBannerConfiguration/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1023','BonusRecords','BonusRecords','',9977,'cfgLst/BonusRecords/index','cfgLst/BonusRecords/index','1001',2,'(*)','cfgLst/BonusRecords/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1024','VIPMonthBonusRecords','VIPMonthBonusRecords','',9976,'cfgLst/VIPMonthBonusRecords/index','cfgLst/VIPMonthBonusRecords/index','1001',2,'(*)','cfgLst/VIPMonthBonusRecords/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1025','VIPWeekBonusRecords','VIPWeekBonusRecords','',9975,'cfgLst/VIPWeekBonusRecords/index','cfgLst/VIPWeekBonusRecords/index','1001',2,'(*)','cfgLst/VIPWeekBonusRecords/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1026','VersionManagement','VersionManagement','',9974,'cfgLst/VersionManagement/index','cfgLst/VersionManagement/index','1001',2,'(*)','cfgLst/VersionManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1027','appLoadingInterface','appLoadingInterface','',9973,'cfgLst/appLoadingInterface/index','cfgLst/appLoadingInterface/index','1001',2,'(*)','cfgLst/appLoadingInterface/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1028','artProxyMgr','artProxyMgr','',9972,'cfgLst/artProxyMgr/index','cfgLst/artProxyMgr/index','1001',2,'(*)','cfgLst/artProxyMgr/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1029','navConfig','navConfig','',9971,'cfgLst/navConfig/index','cfgLst/navConfig/index','1001',2,'(*)','cfgLst/navConfig/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1030','vipUpgRec','vipUpgRec','',9970,'cfgLst/vipUpgRec/index','cfgLst/vipUpgRec/index','1001',2,'(*)','cfgLst/vipUpgRec/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1031','washCode','washCode','',9969,'cfgLst/washCode/index','cfgLst/washCode/index','1001',2,'(*)','cfgLst/washCode/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1033','customerServiceDomains','customerServiceDomains','',9967,'domain/customerServiceDomains/index','domain/customerServiceDomains/index','1002',2,'(*)','domain/customerServiceDomains/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1099','siteLinkManage','siteLinkManage','',9966,'domain/siteLinkManage/index','domain/siteLinkManage/index','1002',2,'(*)','domain/siteLinkManage/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1034','capitalFlow','capitalFlow','',9966,'finance/capitalFlow/index','finance/capitalFlow/index','1003',2,'(*)','finance/capitalFlow/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1035','manualDepositWithdrawal','manualDepositWithdrawal','',9965,'finance/manualDepositWithdrawal/index','finance/manualDepositWithdrawal/index','1003',2,'(*)','finance/manualDepositWithdrawal/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1036','playerTopUp','playerTopUp','',9964,'finance/playerTopUp/index','finance/playerTopUp/index','1003',2,'(*)','finance/playerTopUp/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1037','playerWithDrawal','playerWithDrawal','',9963,'finance/playerWithDrawal/index','finance/playerWithDrawal/index','1003',2,'(*)','finance/playerWithDrawal/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1038','remittanceApplication','remittanceApplication','',9962,'finance/remittanceApplication/index','finance/remittanceApplication/index','1003',2,'(*)','finance/remittanceApplication/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1039','withdrawal','withdrawal','',9961,'finance/withdrawal/index','finance/withdrawal/index','1003',2,'(*)','finance/withdrawal/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1040','withdrawalMessage','withdrawalMessage','',9960,'finance/withdrawalMessage/index','finance/withdrawalMessage/index','1003',2,'(*)','finance/withdrawalMessage/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1041','gameList','gameList','',9959,'gameManage/gameList/index','gameManage/gameList/index','1004',2,'(*)','gameManage/gameList/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1042','gameStadium','gameStadium','',9958,'gameManage/gameStadium/index','gameManage/gameStadium/index','1004',2,'(*)','gameManage/gameStadium/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1043','mallList','mallList','',9957,'gameManage/mallList/index','gameManage/mallList/index','1004',2,'(*)','gameManage/mallList/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1044','merchantManage','merchantManage','',9956,'gameManage/merchantManage/index','gameManage/merchantManage/index','1004',2,'(*)','gameManage/merchantManage/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1045','ticker','ticker','',9955,'gameManage/ticker/index','gameManage/ticker/index','1004',2,'(*)','gameManage/ticker/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1046','IPBlacklistManagement','IPBlacklistManagement','',9954,'home/IPBlacklistManagement/index','home/IPBlacklistManagement/index','1005',2,'(*)','home/IPBlacklistManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1047','IPWhitelistManagement','IPWhitelistManagement','',9953,'home/IPWhitelistManagement/index','home/IPWhitelistManagement/index','1005',2,'(*)','home/IPWhitelistManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1048','accountInfo','accountInfo','',9952,'home/accountInfo/index','home/accountInfo/index','1005',2,'(*)','home/accountInfo/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1049','accountSafe','accountSafe','',9951,'home/accountSafe/index','home/accountSafe/index','1005',2,'(*)','home/accountSafe/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1050','productStatistics','productStatistics','',9950,'home/productStatistics/index','home/productStatistics/index','1005',2,'(*)','home/productStatistics/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1051','alarmMonitoring','alarmMonitoring','',9949,'log/alarmMonitoring/index','log/alarmMonitoring/index','1006',2,'(*)','log/alarmMonitoring/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1052','userOperation','userOperation','',9948,'log/userOperation/index','log/userOperation/index','1006',2,'(*)','log/userOperation/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1053','merchantPlayerQuery','merchantPlayerQuery','',9947,'merchantChannel/merchantPlayerQuery/index','merchantChannel/merchantPlayerQuery/index','1007',2,'(*)','merchantChannel/merchantPlayerQuery/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1054','playerDupeCheck','playerDupeCheck','',9946,'merchantChannel/playerDupeCheck/index','merchantChannel/playerDupeCheck/index','1007',2,'(*)','merchantChannel/playerDupeCheck/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1055','bankCard','bankCard','',9945,'payment/bankCard/index','payment/bankCard/index','1008',2,'(*)','payment/bankCard/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1056','channel','channel','',9944,'payment/channel/index','payment/channel/index','1008',2,'(*)','payment/channel/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1057','channelManagement','channelManagement','',9943,'payment/channelManagement/index','payment/channelManagement/index','1008',2,'(*)','payment/channelManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1058','merchantList','merchantList','',9942,'payment/merchantList/index','payment/merchantList/index','1008',2,'(*)','payment/merchantList/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1059','merchantManagement','merchantManagement','',9941,'payment/merchantManagement/index','payment/merchantManagement/index','1008',2,'(*)','payment/merchantManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1060','merchantWithdrawal','merchantWithdrawal','',9940,'payment/merchantWithdrawal/index','payment/merchantWithdrawal/index','1008',2,'(*)','payment/merchantWithdrawal/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1061','paymentChannelManagement','paymentChannelManagement','',9939,'payment/paymentChannelManagement/index','payment/paymentChannelManagement/index','1008',2,'(*)','payment/paymentChannelManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1062','paymentGatewayManagement','paymentGatewayManagement','',9938,'payment/paymentGatewayManagement/index','payment/paymentGatewayManagement/index','1008',2,'(*)','payment/paymentGatewayManagement/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1063','admin','admin','',9937,'permission/admin/index','permission/admin/index','1009',2,'(*)','permission/admin/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1064','channelPackage','channelPackage','',9936,'permission/channelPackage/index','permission/channelPackage/index','1009',2,'(*)','permission/channelPackage/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1065','menu','menu','',9935,'permission/menu/index','permission/menu/index','1009',2,'(*)','permission/menu/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1066','subchannel','subchannel','',9934,'permission/subchannel/index','permission/subchannel/index','1009',2,'(*)','permission/subchannel/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1067','proxyDomMgmt','proxyDomMgmt','',9933,'proxy/proxyDomMgmt/index','proxy/proxyDomMgmt/index','1010',2,'(*)','proxy/proxyDomMgmt/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1068','proxyList','proxyList','',9932,'proxy/proxyList/index','proxy/proxyList/index','1010',2,'(*)','proxy/proxyList/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1069','LTV','LTV','',9931,'statistics/LTV/index','statistics/LTV/index','1011',2,'(*)','statistics/LTV/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1070','VipUser','VipUser','',9930,'statistics/VipUser/index','statistics/VipUser/index','1011',2,'(*)','statistics/VipUser/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1071','appInstallTracking','appInstallTracking','',9929,'statistics/appInstallTracking/index','statistics/appInstallTracking/index','1011',2,'(*)','statistics/appInstallTracking/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1072','landingPage','landingPage','',9928,'statistics/landingPage/index','statistics/landingPage/index','1011',2,'(*)','statistics/landingPage/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1073','manualReplenishment','manualReplenishment','',9927,'statistics/manualReplenishment/index','statistics/manualReplenishment/index','1011',2,'(*)','statistics/manualReplenishment/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1074','orderQuery','orderQuery','',9926,'statistics/orderQuery/index','statistics/orderQuery/index','1011',2,'(*)','statistics/orderQuery/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1075','playerAccountsReport','playerAccountsReport','',9925,'statistics/playerAccountsReport/index','statistics/playerAccountsReport/index','1011',2,'(*)','statistics/playerAccountsReport/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1076','playerGrouping','playerGrouping','',9924,'statistics/playerGrouping/index','statistics/playerGrouping/index','1011',2,'(*)','statistics/playerGrouping/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1077','promotionDetails','promotionDetails','',9923,'statistics/promotionDetails/index','statistics/promotionDetails/index','1011',2,'(*)','statistics/promotionDetails/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1078','promotionIDReport','promotionIDReport','',9922,'statistics/promotionIDReport/index','statistics/promotionIDReport/index','1011',2,'(*)','statistics/promotionIDReport/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1079','promotionReport','promotionReport','',9921,'statistics/promotionReport/index','statistics/promotionReport/index','1011',2,'(*)','statistics/promotionReport/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1080','reconciliationReport','reconciliationReport','',9920,'statistics/reconciliationReport/index','statistics/reconciliationReport/index','1011',2,'(*)','statistics/reconciliationReport/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1081','retention','retention','',9919,'statistics/retention/index','statistics/retention/index','1011',2,'(*)','statistics/retention/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1082','subgameSummary','subgameSummary','',9918,'statistics/subgameSummary/index','statistics/subgameSummary/index','1011',2,'(*)','statistics/subgameSummary/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1083','summaryReport','summaryReport','',9917,'statistics/summaryReport/index','statistics/summaryReport/index','1011',2,'(*)','statistics/summaryReport/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1084','topGamesChart','topGamesChart','',9916,'statistics/topGamesChart/index','statistics/topGamesChart/index','1011',2,'(*)','statistics/topGamesChart/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1085','userDailyReport','userDailyReport','',9915,'statistics/userDailyReport/index','statistics/userDailyReport/index','1011',2,'(*)','statistics/userDailyReport/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1086','department','department','',9914,'systemManagement/department/index','systemManagement/department/index','1012',2,'(*)','systemManagement/department/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1087','dictTypes','dictTypes','',9913,'systemManagement/dictTypes/index','systemManagement/dictTypes/index','1012',2,'(*)','systemManagement/dictTypes/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1088','dicts','dicts','',9912,'systemManagement/dicts/index','systemManagement/dicts/index','1012',2,'(*)','systemManagement/dicts/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1089','menus','menus','',9911,'systemManagement/menus/index','systemManagement/menus/index','1012',2,'(*)','systemManagement/menus/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1090','role','role','',9910,'systemManagement/role/index','systemManagement/role/index','1012',2,'(*)','systemManagement/role/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1091','systemUser','systemUser','',9909,'systemManagement/systemUser/index','systemManagement/systemUser/index','1012',2,'(*)','systemManagement/systemUser/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1092','labels','labels','',9908,'userLabels/labels/index','userLabels/labels/index','1013',2,'(*)','userLabels/labels/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1093','labelsSeach','labelsSeach','',9907,'userLabels/labelsSeach/index','userLabels/labelsSeach/index','1013',2,'(*)','userLabels/labelsSeach/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1094','telCode','telCode','',9906,'userManage/telCode/index','userManage/telCode/index','1014',2,'(*)','userManage/telCode/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1095','transferAccounts','transferAccounts','',9905,'userManage/transferAccounts/index','userManage/transferAccounts/index','1014',2,'(*)','userManage/transferAccounts/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1096','user','user','',9904,'userManage/user/index','userManage/user/index','1014',2,'(*)','userManage/user/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1097','vip','vip','',9903,'userManage/vip/index','userManage/vip/index','1014',2,'(*)','userManage/vip/index',1,0,'');
INSERT INTO `menu` (`id`,`name`,`menu_name`,`icon`,`sort`,`role_flag`,`address`,`parent_id`,`type`,`api_regular`,`perms`,`show_status`,`open_cache`,`locales`) VALUES ('1098','vipInfo','vipInfo','',9902,'userManage/vipInfo/index','userManage/vipInfo/index','1014',2,'(*)','userManage/vipInfo/index',1,0,'');
UPDATE `role` SET `meus_ids` = '1000,1001,1002,1003,1004,1005,1006,1007,1008,1009,1010,1011,1012,1013,1014,1015,1016,1017,1018,1019,1020,1021,1022,1023,1024,1025,1026,1027,1028,1029,1030,1031,1032,1033,1034,1035,1036,1037,1038,1039,1040,1041,1042,1043,1044,1045,1046,1047,1048,1049,1050,1051,1052,1053,1054,1055,1056,1057,1058,1059,1060,1061,1062,1063,1064,1065,1066,1067,1068,1069,1070,1071,1072,1073,1074,1075,1076,1077,1078,1079,1080,1081,1082,1083,1084,1085,1086,1087,1088,1089,1090,1091,1092,1093,1094,1095,1096,1097,1098,1099', `perms_list` = '["*"]' WHERE `id` = '1';

-- Keep local seed output in zh-cn names for admin UI
UPDATE menu SET menu_name='活动管理', name='活动管理' WHERE role_flag='activity';
UPDATE menu SET menu_name='配置管理', name='配置管理' WHERE role_flag='cfgLst';
UPDATE menu SET menu_name='域名管理', name='域名管理' WHERE role_flag='domain';
UPDATE menu SET menu_name='财务管理', name='财务管理' WHERE role_flag='finance';
UPDATE menu SET menu_name='游戏管理', name='游戏管理' WHERE role_flag='gameManage';
UPDATE menu SET menu_name='首页管理', name='首页管理' WHERE role_flag='home';
UPDATE menu SET menu_name='日志管理', name='日志管理' WHERE role_flag='log';
UPDATE menu SET menu_name='渠道管理', name='渠道管理' WHERE role_flag='merchantChannel';
UPDATE menu SET menu_name='支付管理', name='支付管理' WHERE role_flag='payment';
UPDATE menu SET menu_name='权限管理', name='权限管理' WHERE role_flag='permission';
UPDATE menu SET menu_name='代理管理', name='代理管理' WHERE role_flag='proxy';
UPDATE menu SET menu_name='统计报表', name='统计报表' WHERE role_flag='statistics';
UPDATE menu SET menu_name='系统管理', name='系统管理' WHERE role_flag='systemManagement';
UPDATE menu SET menu_name='用户标签', name='用户标签' WHERE role_flag='userLabels';
UPDATE menu SET menu_name='用户管理', name='用户管理' WHERE role_flag='userManage';
UPDATE menu SET menu_name='优惠活动', name='优惠活动' WHERE role_flag LIKE '%/benefit/index';
UPDATE menu SET menu_name='每日签到', name='每日签到' WHERE role_flag LIKE '%/checkIn/index';
UPDATE menu SET menu_name='新活动', name='新活动' WHERE role_flag LIKE '%/newActivity/index';
UPDATE menu SET menu_name='公告管理', name='公告管理' WHERE role_flag LIKE '%/notice/index';
UPDATE menu SET menu_name='短信管理', name='短信管理' WHERE role_flag LIKE '%/shortMessageManagement/index';
UPDATE menu SET menu_name='任务管理', name='任务管理' WHERE role_flag LIKE '%/task/index';
UPDATE menu SET menu_name='卡券管理', name='卡券管理' WHERE role_flag LIKE '%/ticket/index';
UPDATE menu SET menu_name='广告配置', name='广告配置' WHERE role_flag LIKE '%/AdBannerConfiguration/index';
UPDATE menu SET menu_name='红利记录', name='红利记录' WHERE role_flag LIKE '%/BonusRecords/index';
UPDATE menu SET menu_name='VIP月礼记录', name='VIP月礼记录' WHERE role_flag LIKE '%/VIPMonthBonusRecords/index';
UPDATE menu SET menu_name='VIP周礼记录', name='VIP周礼记录' WHERE role_flag LIKE '%/VIPWeekBonusRecords/index';
UPDATE menu SET menu_name='版本管理', name='版本管理' WHERE role_flag LIKE '%/VersionManagement/index';
UPDATE menu SET menu_name='启动页配置', name='启动页配置' WHERE role_flag LIKE '%/appLoadingInterface/index';
UPDATE menu SET menu_name='艺术代理管理', name='艺术代理管理' WHERE role_flag LIKE '%/artProxyMgr/index';
UPDATE menu SET menu_name='导航配置', name='导航配置' WHERE role_flag LIKE '%/navConfig/index';
UPDATE menu SET menu_name='VIP升级记录', name='VIP升级记录' WHERE role_flag LIKE '%/vipUpgRec/index';
UPDATE menu SET menu_name='洗码配置', name='洗码配置' WHERE role_flag LIKE '%/washCode/index';
UPDATE menu SET menu_name='客服域名', name='客服域名' WHERE role_flag LIKE '%/customerServiceDomains/index';
UPDATE menu SET menu_name='站点链接配置', name='站点链接配置' WHERE role_flag LIKE '%/siteLinkManage/index';
UPDATE menu SET menu_name='资金流水', name='资金流水' WHERE role_flag LIKE '%/capitalFlow/index';
UPDATE menu SET menu_name='人工存提', name='人工存提' WHERE role_flag LIKE '%/manualDepositWithdrawal/index';
UPDATE menu SET menu_name='玩家充值', name='玩家充值' WHERE role_flag LIKE '%/playerTopUp/index';
UPDATE menu SET menu_name='玩家提现', name='玩家提现' WHERE role_flag LIKE '%/playerWithDrawal/index';
UPDATE menu SET menu_name='汇款申请', name='汇款申请' WHERE role_flag LIKE '%/remittanceApplication/index';
UPDATE menu SET menu_name='提现管理', name='提现管理' WHERE role_flag LIKE '%/withdrawal/index';
UPDATE menu SET menu_name='提现信息维护', name='提现信息维护' WHERE role_flag LIKE '%/withdrawalMessage/index';
UPDATE menu SET menu_name='游戏列表', name='游戏列表' WHERE role_flag LIKE '%/gameList/index';
UPDATE menu SET menu_name='游戏场馆', name='游戏场馆' WHERE role_flag LIKE '%/gameStadium/index';
UPDATE menu SET menu_name='商城列表', name='商城列表' WHERE role_flag LIKE '%/mallList/index';
UPDATE menu SET menu_name='商户管理', name='商户管理' WHERE role_flag LIKE '%/merchantManage/index';
UPDATE menu SET menu_name='跑马灯管理', name='跑马灯管理' WHERE role_flag LIKE '%/ticker/index';
UPDATE menu SET menu_name='IP黑名单', name='IP黑名单' WHERE role_flag LIKE '%/IPBlacklistManagement/index';
UPDATE menu SET menu_name='IP白名单', name='IP白名单' WHERE role_flag LIKE '%/IPWhitelistManagement/index';
UPDATE menu SET menu_name='账号信息', name='账号信息' WHERE role_flag LIKE '%/accountInfo/index';
UPDATE menu SET menu_name='账号安全', name='账号安全' WHERE role_flag LIKE '%/accountSafe/index';
UPDATE menu SET menu_name='产品统计', name='产品统计' WHERE role_flag LIKE '%/productStatistics/index';
UPDATE menu SET menu_name='告警监控', name='告警监控' WHERE role_flag LIKE '%/alarmMonitoring/index';
UPDATE menu SET menu_name='用户操作日志', name='用户操作日志' WHERE role_flag LIKE '%/userOperation/index';
UPDATE menu SET menu_name='商户玩家查询', name='商户玩家查询' WHERE role_flag LIKE '%/merchantPlayerQuery/index';
UPDATE menu SET menu_name='玩家查重', name='玩家查重' WHERE role_flag LIKE '%/playerDupeCheck/index';
UPDATE menu SET menu_name='银行卡管理', name='银行卡管理' WHERE role_flag LIKE '%/bankCard/index';
UPDATE menu SET menu_name='支付通道', name='支付通道' WHERE role_flag LIKE '%/channel/index';
UPDATE menu SET menu_name='渠道管理', name='渠道管理' WHERE role_flag LIKE '%/channelManagement/index';
UPDATE menu SET menu_name='商户列表', name='商户列表' WHERE role_flag LIKE '%/merchantList/index';
UPDATE menu SET menu_name='商户管理', name='商户管理' WHERE role_flag LIKE '%/merchantManagement/index';
UPDATE menu SET menu_name='商户提现', name='商户提现' WHERE role_flag LIKE '%/merchantWithdrawal/index';
UPDATE menu SET menu_name='支付通道管理', name='支付通道管理' WHERE role_flag LIKE '%/paymentChannelManagement/index';
UPDATE menu SET menu_name='支付网关管理', name='支付网关管理' WHERE role_flag LIKE '%/paymentGatewayManagement/index';
UPDATE menu SET menu_name='管理员管理', name='管理员管理' WHERE role_flag LIKE '%/admin/index';
UPDATE menu SET menu_name='渠道包管理', name='渠道包管理' WHERE role_flag LIKE '%/channelPackage/index';
UPDATE menu SET menu_name='菜单权限', name='菜单权限' WHERE role_flag LIKE '%/menu/index';
UPDATE menu SET menu_name='子渠道管理', name='子渠道管理' WHERE role_flag LIKE '%/subchannel/index';
UPDATE menu SET menu_name='代理域名管理', name='代理域名管理' WHERE role_flag LIKE '%/proxyDomMgmt/index';
UPDATE menu SET menu_name='代理列表', name='代理列表' WHERE role_flag LIKE '%/proxyList/index';
UPDATE menu SET menu_name='LTV分析', name='LTV分析' WHERE role_flag LIKE '%/LTV/index';
UPDATE menu SET menu_name='VIP用户分析', name='VIP用户分析' WHERE role_flag LIKE '%/VipUser/index';
UPDATE menu SET menu_name='安装追踪', name='安装追踪' WHERE role_flag LIKE '%/appInstallTracking/index';
UPDATE menu SET menu_name='落地页统计', name='落地页统计' WHERE role_flag LIKE '%/landingPage/index';
UPDATE menu SET menu_name='人工补单', name='人工补单' WHERE role_flag LIKE '%/manualReplenishment/index';
UPDATE menu SET menu_name='订单查询', name='订单查询' WHERE role_flag LIKE '%/orderQuery/index';
UPDATE menu SET menu_name='玩家账务报表', name='玩家账务报表' WHERE role_flag LIKE '%/playerAccountsReport/index';
UPDATE menu SET menu_name='玩家分组', name='玩家分组' WHERE role_flag LIKE '%/playerGrouping/index';
UPDATE menu SET menu_name='推广明细', name='推广明细' WHERE role_flag LIKE '%/promotionDetails/index';
UPDATE menu SET menu_name='推广ID报表', name='推广ID报表' WHERE role_flag LIKE '%/promotionIDReport/index';
UPDATE menu SET menu_name='推广报表', name='推广报表' WHERE role_flag LIKE '%/promotionReport/index';
UPDATE menu SET menu_name='对账报表', name='对账报表' WHERE role_flag LIKE '%/reconciliationReport/index';
UPDATE menu SET menu_name='留存分析', name='留存分析' WHERE role_flag LIKE '%/retention/index';
UPDATE menu SET menu_name='子游戏汇总', name='子游戏汇总' WHERE role_flag LIKE '%/subgameSummary/index';
UPDATE menu SET menu_name='汇总报表', name='汇总报表' WHERE role_flag LIKE '%/summaryReport/index';
UPDATE menu SET menu_name='热门游戏排行', name='热门游戏排行' WHERE role_flag LIKE '%/topGamesChart/index';
UPDATE menu SET menu_name='用户日报', name='用户日报' WHERE role_flag LIKE '%/userDailyReport/index';
UPDATE menu SET menu_name='部门管理', name='部门管理' WHERE role_flag LIKE '%/department/index';
UPDATE menu SET menu_name='字典类型', name='字典类型' WHERE role_flag LIKE '%/dictTypes/index';
UPDATE menu SET menu_name='字典数据', name='字典数据' WHERE role_flag LIKE '%/dicts/index';
UPDATE menu SET menu_name='菜单管理', name='菜单管理' WHERE role_flag LIKE '%/menus/index';
UPDATE menu SET menu_name='角色管理', name='角色管理' WHERE role_flag LIKE '%/role/index';
UPDATE menu SET menu_name='系统用户', name='系统用户' WHERE role_flag LIKE '%/systemUser/index';
UPDATE menu SET menu_name='标签管理', name='标签管理' WHERE role_flag LIKE '%/labels/index';
UPDATE menu SET menu_name='标签查询', name='标签查询' WHERE role_flag LIKE '%/labelsSeach/index';
UPDATE menu SET menu_name='手机验证码', name='手机验证码' WHERE role_flag LIKE '%/telCode/index';
UPDATE menu SET menu_name='转账管理', name='转账管理' WHERE role_flag LIKE '%/transferAccounts/index';
UPDATE menu SET menu_name='用户管理', name='用户管理' WHERE role_flag LIKE '%/user/index';
UPDATE menu SET menu_name='VIP管理', name='VIP管理' WHERE role_flag LIKE '%/vip/index';
UPDATE menu SET menu_name='VIP信息', name='VIP信息' WHERE role_flag LIKE '%/vipInfo/index';

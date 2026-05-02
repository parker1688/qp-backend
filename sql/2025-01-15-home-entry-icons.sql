/**
 * 表fc_merchant_link新增字段home_entry_icons - 首页入口图标配置
 * 日期: 2025-01-01
 */
ALTER TABLE fc_merchant_link 
ADD COLUMN home_entry_icons LONGTEXT COMMENT '首页入口图标配置 JSON (格式: {homeEntryIcons: [{name, code, image, imageOK, venueType}, ...]})';

/**
 * 初始化商户的首页入口图标配置（示例数据）
 */
UPDATE fc_merchant_link 
SET home_entry_icons = '{
  "homeEntryIcons": [
    {
      "name": "热门",
      "code": "rm",
      "index": 0,
      "image": "/static/home/menu/icon_rm.png",
      "imageOK": "/static/home/menu/icon_rm.png",
      "venueType": "POPULAR"
    },
    {
      "name": "电子",
      "code": "dz",
      "index": 1,
      "image": "/static/home/menu/icon_dz.png",
      "imageOK": "/static/home/menu/icon_dz.png",
      "venueType": "DZ"
    },
    {
      "name": "棋牌",
      "code": "qp",
      "index": 2,
      "image": "/static/home/menu/icon_qp.png",
      "imageOK": "/static/home/menu/icon_qp.png",
      "venueType": "QP"
    },
    {
      "name": "捕鱼",
      "code": "by",
      "index": 3,
      "image": "/static/home/menu/icon_by.png",
      "imageOK": "/static/home/menu/icon_by.png",
      "venueType": "BY"
    },
    {
      "name": "真人",
      "code": "zr",
      "index": 4,
      "image": "/static/home/menu/icon_zr.png",
      "imageOK": "/static/home/menu/icon_zr.png",
      "venueType": "ZR"
    },
    {
      "name": "体育",
      "code": "ty",
      "index": 5,
      "image": "/static/home/menu/icon_ty.png",
      "imageOK": "/static/home/menu/icon_ty.png",
      "venueType": "TY"
    },
    {
      "name": "彩票",
      "code": "cp",
      "index": 6,
      "image": "/static/home/menu/icon_cp.png",
      "imageOK": "/static/home/menu/icon_cp.png",
      "venueType": "CP"
    }
  ]
}'
WHERE id IS NOT NULL;

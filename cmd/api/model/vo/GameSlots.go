package vo

type GameSlotsResp struct {
	Id             string `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	VenueCode      string `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                         // 场馆code
	VenueType      string `gorm:"venue_type" json:"venue_type" form:"venue_type" uri:"venue_type" `                         // 场馆类型
	GameCode       string `gorm:"game_code" json:"game_code" form:"game_code" uri:"game_code" `                             // 游戏Code
	Hot            int    `gorm:"hot" json:"hot" form:"hot" uri:"hot" `                                                     // 是否热门 0: 不热门 1: 热门
	Sort           int    `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                                 // 排序
	ImgIcon        string `gorm:"img_icon" json:"img_icon" form:"img_icon" uri:"img_icon" `                                 // 图片地址
	GameName       string `gorm:"game_name" json:"game_name" form:"game_name" uri:"game_name" `                             // 游戏名称
	GameNamePinyin string `gorm:"game_name_pinyin" json:"game_name_pinyin" form:"game_name_pinyin" uri:"game_name_pinyin" ` // 游戏名称英文
	GameType       string `gorm:"game_type" json:"game_type" form:"game_type" uri:"game_type" `                             // 游戏分类标识
	Status         int    `gorm:"status" json:"status" form:"status" uri:"status" `                                         // 状态  1:上线  2:下线  3: 维护
	Gtype          *int   `gorm:"gtype;default:null" json:"gtype" form:"gtype" uri:"gtype" `                                // 游戏种类
}

type HistoryGameSlotsReq struct {
	Id             string `gorm:"id;primary_key;AUTO_INCREMENT" json:"id" form:"id" uri:"id" `
	VenueCode      string `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                         // 场馆code
	VenueType      string `gorm:"venue_type" json:"venue_type" form:"venue_type" uri:"venue_type" `                         // 场馆类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	GameCode       string `gorm:"game_code" json:"game_code" form:"game_code" uri:"game_code" `                             // 游戏Code
	ImgIcon        string `gorm:"img_icon" json:"img_icon" form:"img_icon" uri:"img_icon" `                                 // 图片地址
	GameName       string `gorm:"game_name" json:"game_name" form:"game_name" uri:"game_name" `                             // 游戏名称
	GameNamePinyin string `gorm:"game_name_pinyin" json:"game_name_pinyin" form:"game_name_pinyin" uri:"game_name_pinyin" ` // 游戏名称英文
}

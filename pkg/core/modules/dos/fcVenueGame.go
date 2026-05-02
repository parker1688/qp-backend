package dos

import (
	"bootpkg/common/expands/automaticType"
)

type FcVenueGame struct {
	BaseDos
	VenueName      string             `gorm:"venue_name" json:"venue_name" form:"venue_name" uri:"venue_name" `                         // 场馆名字
	VenueCode      string             `gorm:"venue_code" json:"venue_code" form:"venue_code" uri:"venue_code" `                         // 场馆code
	VenueType      string             `gorm:"venue_type" json:"venue_type" form:"venue_type" uri:"venue_type" `                         // 场馆类型 场馆类型 chess 棋牌,elecgame 电游,live 真人,sport 体育,esport 电竞,lottery 彩票,fish 捕鱼
	Status         int                `gorm:"status" json:"status" form:"status" uri:"status" `                                         // 状态  1:上线  2:下线  3: 维护
	CreateTime     automaticType.Time `gorm:"create_time;default:null" json:"create_time" form:"create_time" uri:"create_time" `        // 创建时间
	CreateBy       string             `gorm:"create_by" json:"-" form:"create_by" uri:"create_by" `                                     // 创建人
	UpdateTime     automaticType.Time `gorm:"update_time;default:null" json:"update_time" form:"update_time" uri:"update_time" `        // 修改时间
	UpdateBy       string             `gorm:"update_by" json:"-" form:"update_by" uri:"update_by" `                                     // 修改人
	GameCode       string             `gorm:"game_code" json:"game_code" form:"game_code" uri:"game_code" `                             // 游戏Code
	Hot            int                `gorm:"hot" json:"hot" form:"hot" uri:"hot" `                                                     // 是否热门 0: 不热门 1: 热门
	Sort           int                `gorm:"sort" json:"sort" form:"sort" uri:"sort" `                                                 // 排序
	ImgIcon        string             `gorm:"img_icon" json:"img_icon" form:"img_icon" uri:"img_icon" `                                 // 图片地址
	GameName       string             `gorm:"game_name" json:"game_name" form:"game_name" uri:"game_name" `                             // 游戏名称
	GameNamePinyin string             `gorm:"game_name_pinyin" json:"game_name_pinyin" form:"game_name_pinyin" uri:"game_name_pinyin" ` // 游戏名称英文
	Remark         string             `gorm:"remark" json:"remark" form:"remark" uri:"remark" `                                         // 备注
	Language       string             `gorm:"language" json:"language" form:"language" uri:"language" `                                 // 语言简码
	GameType       string             `gorm:"game_type" json:"game_type" form:"game_type" uri:"game_type" `                             // 游戏分类标识
	Gtype          *int               `gorm:"gtype;default:null" json:"gtype" form:"gtype" uri:"gtype" `                                // 游戏种类
}

func (FcVenueGame) TableName() string {
	return "fc_venue_game"
}

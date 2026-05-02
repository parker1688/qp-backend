package venues

import (
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/enmus"
	"bootpkg/pkg/service/venues/venueDetail"
)

// venuesLineConfig
//
//	@Description: 场馆配置线路
//	@param lineName 商户名称
//	@param venueCode 场馆Code
//	@return venueDetail.IVenues 实际场馆
//	@return int 当前商户配置的线路
func venuesLineConfig(merchantCode string, venueCode string) (venueDetail.IVenues, int) {
	//var lineNum = 0
	var lineNum = 1 //先全部为1  后期要实现这个逻辑
	if len(merchantCode) > 0 {
		//根据商户Code获取线路配置
	}
	switch venueCode {
	case enmus.FBTY:
		return venueDetail.NewVenueFBTY(&global.CONFIG.Venue.Fbty), lineNum
	case enmus.HGTY:
		return venueDetail.NewHGTY(&global.CONFIG.Venue.HGTY), lineNum
	case enmus.WUGDZ:
		return venueDetail.NewWUGDZ(&global.CONFIG.Venue.WUGDZ), lineNum
	case enmus.FGDZ:
		return venueDetail.NewFGDZ(&global.CONFIG.Venue.FGDZ), lineNum
	case enmus.DBDJ:
		return venueDetail.NewDBDJ(&global.CONFIG.Venue.DBDJ), lineNum
	case enmus.KYQP:
		return venueDetail.NewKYQP(&global.CONFIG.Venue.KYQP), lineNum
	case enmus.TYQP:
		return venueDetail.NewTYQP(&global.CONFIG.Venue.TYQP), lineNum
	case enmus.LYQP:
		return venueDetail.NewLYQP(&global.CONFIG.Venue.LYQP), lineNum
	case enmus.VGQP:
		return venueDetail.NewVGQP(&global.CONFIG.Venue.VGQP), lineNum
	case enmus.MTQP:
		return venueDetail.NewVenueMTQP(&global.CONFIG.Venue.MTQP), lineNum
	case enmus.MGDZ:
		return venueDetail.NewVenueMGDZ(&global.CONFIG.Venue.MGDZ), lineNum
	case enmus.PGDZ:
		return venueDetail.NewPGDZ(&global.CONFIG.Venue.PGDZ), lineNum
	case enmus.AGZR:
		return venueDetail.NewAGZR(&global.CONFIG.Venue.AGZR), lineNum
	case enmus.BGZR:
		return venueDetail.NewBGZR(&global.CONFIG.Venue.BGZR), lineNum
	case enmus.DBZR:
		return venueDetail.NewDBZR(&global.CONFIG.Venue.DBZR), lineNum
	case enmus.CQ9:
		return venueDetail.NewCQ9(&global.CONFIG.Venue.CQ9), lineNum
	case enmus.JDB:
		return venueDetail.NewJDB(&global.CONFIG.Venue.JDB), lineNum
	case enmus.VRCP:
		return venueDetail.NewVRCP(&global.CONFIG.Venue.VRCP), lineNum
	case enmus.BBIN:
		return venueDetail.NewBBIN(&global.CONFIG.Venue.BBIN), lineNum
	case enmus.WALI:
		return venueDetail.NewWALI(&global.CONFIG.Venue.WALI), lineNum
	case enmus.PPDZ:
		return venueDetail.NewPPDZ(&global.CONFIG.Venue.PPDZ), lineNum
	case enmus.PTDZ:
		return venueDetail.NewPTDZ(&global.CONFIG.Venue.PTDZ), lineNum
	case enmus.JLDZ:
		return venueDetail.NewJLDZ(&global.CONFIG.Venue.JLDZ), lineNum
	case enmus.KXDZ:
		return venueDetail.NewKXDZ(&global.CONFIG.Venue.KXDZ), lineNum
	case enmus.LGDDZ:
		return venueDetail.NewLGDDZ(&global.CONFIG.Venue.LGDDZ), lineNum
	case enmus.MWDZ:
		return venueDetail.NewMWDZ(&global.CONFIG.Venue.MWDZ), lineNum
	case enmus.SGDZ:
		return venueDetail.NewSGDZ(&global.CONFIG.Venue.SGDZ), lineNum
	case enmus.TTQP:
		return venueDetail.NewTTQP(&global.CONFIG.Venue.TTQP), lineNum
	case enmus.IMTY:
		return venueDetail.NewIMTY(&global.CONFIG.Venue.IMTY), lineNum
	case enmus.SBTY:
		return venueDetail.NewSBTY(&global.CONFIG.Venue.SBTY), lineNum
	case enmus.PDTY:
		return venueDetail.NewPDTY(&global.CONFIG.Venue.PDTY), lineNum
	}

	return nil, -1

}

package enmus

const (
	FBTY  = "FBTY"
	HGTY  = "HGTY"
	PGDZ  = "PGDZ"
	WUGDZ = "WUGDZ"
	FGDZ  = "FGDZ"
	DBDJ  = "DBDJ"
	KYQP  = "KYQP"
	TYQP  = "TYQP" //天游棋牌,天豪棋牌
	VGQP  = "VGQP"
	LYQP  = "LYQP"
	MTQP  = "MTQP"
	MGDZ  = "MGDZ"
	AGZR  = "AGZR"
	BGZR  = "BGZR"
	JDB   = "JDB"
	CQ9   = "CQ9"
	VRCP  = "VRCP"
	DBZR  = "DBZR"
	BBIN  = "BBIN"
	WALI  = "WALI"
	PPDZ  = "PPDZ"
	PTDZ  = "PTDZ"
	JLDZ  = "JLDZ"
	KXDZ  = "KXDZ"
	LGDDZ = "LGDDZ"
	MWDZ  = "MWDZ"
	SGDZ  = "SGDZ"
	TTQP  = "TTQP"
	IMTY  = "IMTY"
	SBTY  = "SBTY"
	PDTY  = "PDTY"
)

const (
	TransferStatusSuccess    = "success"    //转账成功
	TransferStatusFail       = "fail"       //转账失败
	TransferStatusProcessing = "processing" //转账处理中
)

const (
	//1 体育 2 真人 3 棋牌 4 电子 5 捕鱼 6  彩票  7  区块链 8:电竞

	Game_type_sport      = 1
	Game_type_live       = 2
	Game_type_chess      = 3
	Game_type_elec       = 4
	Game_type_fish       = 5
	Game_type_lottery    = 6
	Game_type_blockchain = 7
	Game_type_esport     = 8
	Game_type_slots      = 9  //老虎机游戏类型。
	Game_type_table      = 10 //赌桌游戏类型。
	Game_type_arcade     = 11 //街机游戏类型。
	Game_type_poker      = 12 //扑克游戏类型。
	Game_type_bingo      = 13 //
)

const (
	Venue_Type_Wallet_Transfer = 1 //转账钱包
	Venue_Type_Wallet_Seamless = 2 //单一钱包
)

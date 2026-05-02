// The build tag makes sure the stub is not built in the final build.

package modules

import (
	"bootpkg/common/ecode"
	"bootpkg/common/global"
	"bootpkg/pkg/core/modules/dos"
)

func SaveFcBindBlockchainType(vo *dos.FcBindBlockchainType) (bool, ecode.Code) {
	err := global.G_DB.Create(vo).Error
	if err != nil {
		return false, ecode.FAIL
	}
	return true, ecode.OK
}

func FindPageFcBindBlockchainType(page, pageSize int, vo *dos.FcBindBlockchainType) (ret []*dos.FcBindBlockchainType, total int64) {
	query := global.G_DB.Model(&dos.FcBindBlockchainType{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BlockchainImg) > 0 {
		query = query.Where("blockchain_img = ?", vo.BlockchainImg)
	}

	if len(vo.ContractName) > 0 {
		query = query.Where("contract_name = ?", vo.ContractName)
	}

	var count int64
	query.Count(&count)
	var dataSlice []*dos.FcBindBlockchainType
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&dataSlice)
	return dataSlice, count
}

func FindByKeyFcBindBlockchainType(vo *dos.FcBindBlockchainType) []*dos.FcBindBlockchainType {
	var data []*dos.FcBindBlockchainType
	query := global.G_DB.Model(&dos.FcBindBlockchainType{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BlockchainImg) > 0 {
		query = query.Where("blockchain_img = ?", vo.BlockchainImg)
	}

	if len(vo.ContractName) > 0 {
		query = query.Where("contract_name = ?", vo.ContractName)
	}

	query.Order("sort desc").Find(&data)
	return data
}

func FindByKeyFcBindBlockchainTypeFirst(vo *dos.FcBindBlockchainType) *dos.FcBindBlockchainType {
	var data *dos.FcBindBlockchainType
	query := global.G_DB.Model(&dos.FcBindBlockchainType{})

	if len(vo.Id) > 0 {
		query = query.Where("id = ?", vo.Id)
	}

	if len(vo.Blockchain) > 0 {
		query = query.Where("blockchain = ?", vo.Blockchain)
	}

	if len(vo.ContractType) > 0 {
		query = query.Where("contract_type = ?", vo.ContractType)
	}

	if vo.Sort > 0 {
		query = query.Where("sort = ?", vo.Sort)
	}

	if len(vo.BlockchainImg) > 0 {
		query = query.Where("blockchain_img = ?", vo.BlockchainImg)
	}

	if len(vo.ContractName) > 0 {
		query = query.Where("contract_name = ?", vo.ContractName)
	}

	query.Take(&data)
	return data
}

// 根据主键Update
func UpdateFcBindBlockchainType(vo *dos.FcBindBlockchainType) bool {
	return global.G_DB.Model(vo).Where(`id = ?`, vo.Id).Updates(map[string]interface{}{
		"blockchain":     vo.Blockchain,
		"contract_type":  vo.ContractType,
		"sort":           vo.Sort,
		"blockchain_img": vo.BlockchainImg,
		"contract_name":  vo.ContractName,
	}).Error == nil
}

func DeleteFcBindBlockchainType(vo *dos.FcBindBlockchainType) bool {
	return global.G_DB.Model(&dos.FcBindBlockchainType{}).Where("id = ?", vo.Id).Delete(vo).Error == nil
}

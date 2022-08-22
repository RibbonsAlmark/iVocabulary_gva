package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
)

type VocabularyInfoService struct {
}

// CreateVocabularyInfo 创建VocabularyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocService *VocabularyInfoService) CreateVocabularyInfo(voc ivocabulary.VocabularyInfo) (err error) {
	id := global.IV_SF.Generate()
	*voc.Uuid = int(id.Int64())
	err = global.GVA_DB.Create(&voc).Error
	// // TMP ...
	// fmt.Printf("Int64  ID: %d\n", id)
	// fmt.Printf("String ID: %s\n", id)
	// fmt.Printf("Base2  ID: %s\n", id.Base2())
	// fmt.Printf("Base64 ID: %s\n", id.Base64())
	// fmt.Printf("ID Time  : %d\n", id.Time())
	// fmt.Printf("ID Node  : %d\n", id.Node())
	// fmt.Printf("ID Step  : %d\n", id.Step())
	// // TMP ...
	return err
}

// DeleteVocabularyInfo 删除VocabularyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocService *VocabularyInfoService) DeleteVocabularyInfo(voc ivocabulary.VocabularyInfo) (err error) {
	err = global.GVA_DB.Delete(&voc).Error
	return err
}

// DeleteVocabularyInfoByIds 批量删除VocabularyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocService *VocabularyInfoService) DeleteVocabularyInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ivocabulary.VocabularyInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateVocabularyInfo 更新VocabularyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocService *VocabularyInfoService) UpdateVocabularyInfo(voc ivocabulary.VocabularyInfo) (err error) {
	err = global.GVA_DB.Save(&voc).Error
	return err
}

// GetVocabularyInfo 根据id获取VocabularyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocService *VocabularyInfoService) GetVocabularyInfo(id uint) (err error, voc ivocabulary.VocabularyInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&voc).Error
	return
}

// GetVocabularyInfoInfoList 分页获取VocabularyInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocService *VocabularyInfoService) GetVocabularyInfoInfoList(info ivocabularyReq.VocabularyInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ivocabulary.VocabularyInfo{})
	var vocs []ivocabulary.VocabularyInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&vocs).Error
	return err, vocs, total
}

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
)

type IvVocabulariesService struct {
}

// CreateIvVocabularies 创建IvVocabularies记录
// Author [piexlmax](https://github.com/piexlmax)
func (ivVocabulariesService *IvVocabulariesService) CreateIvVocabularies(ivVocabularies ivocabulary.IvVocabularies) (err error) {
	err = global.GVA_DB.Create(&ivVocabularies).Error
	return err
}

// DeleteIvVocabularies 删除IvVocabularies记录
// Author [piexlmax](https://github.com/piexlmax)
func (ivVocabulariesService *IvVocabulariesService)DeleteIvVocabularies(ivVocabularies ivocabulary.IvVocabularies) (err error) {
	err = global.GVA_DB.Delete(&ivVocabularies).Error
	return err
}

// DeleteIvVocabulariesByIds 批量删除IvVocabularies记录
// Author [piexlmax](https://github.com/piexlmax)
func (ivVocabulariesService *IvVocabulariesService)DeleteIvVocabulariesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ivocabulary.IvVocabularies{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIvVocabularies 更新IvVocabularies记录
// Author [piexlmax](https://github.com/piexlmax)
func (ivVocabulariesService *IvVocabulariesService)UpdateIvVocabularies(ivVocabularies ivocabulary.IvVocabularies) (err error) {
	err = global.GVA_DB.Save(&ivVocabularies).Error
	return err
}

// GetIvVocabularies 根据id获取IvVocabularies记录
// Author [piexlmax](https://github.com/piexlmax)
func (ivVocabulariesService *IvVocabulariesService)GetIvVocabularies(id uint) (err error, ivVocabularies ivocabulary.IvVocabularies) {
	err = global.GVA_DB.Where("id = ?", id).First(&ivVocabularies).Error
	return
}

// GetIvVocabulariesInfoList 分页获取IvVocabularies记录
// Author [piexlmax](https://github.com/piexlmax)
func (ivVocabulariesService *IvVocabulariesService)GetIvVocabulariesInfoList(info ivocabularyReq.IvVocabulariesSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ivocabulary.IvVocabularies{})
    var ivVocabulariess []ivocabulary.IvVocabularies
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&ivVocabulariess).Error
	return err, ivVocabulariess, total
}

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
)

type WordInfoService struct {
}

// CreateWordInfo 创建WordInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordsService *WordInfoService) CreateWordInfo(words ivocabulary.WordInfo) (err error) {
	*words.Uuid = int(global.IV_SF.Generate().Int64())
	err = global.GVA_DB.Create(&words).Error
	return err
}

// DeleteWordInfo 删除WordInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordsService *WordInfoService) DeleteWordInfo(words ivocabulary.WordInfo) (err error) {
	err = global.GVA_DB.Delete(&words).Error
	return err
}

// DeleteWordInfoByIds 批量删除WordInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordsService *WordInfoService) DeleteWordInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ivocabulary.WordInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateWordInfo 更新WordInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordsService *WordInfoService) UpdateWordInfo(words ivocabulary.WordInfo) (err error) {
	err = global.GVA_DB.Save(&words).Error
	return err
}

// GetWordInfo 根据id获取WordInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordsService *WordInfoService) GetWordInfo(id uint) (err error, words ivocabulary.WordInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&words).Error
	return
}

// GetWordInfoInfoList 分页获取WordInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordsService *WordInfoService) GetWordInfoInfoList(info ivocabularyReq.WordInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&ivocabulary.WordInfo{})
	var wordss []ivocabulary.WordInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&wordss).Error
	return err, wordss, total
}

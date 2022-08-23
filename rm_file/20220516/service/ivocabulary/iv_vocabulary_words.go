package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
)

type VocabularyWordsService struct {
}

// CreateVocabularyWords 创建VocabularyWords记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocWordsService *VocabularyWordsService) CreateVocabularyWords(vocWords ivocabulary.VocabularyWords) (err error) {
	err = global.GVA_DB.Create(&vocWords).Error
	return err
}

// DeleteVocabularyWords 删除VocabularyWords记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocWordsService *VocabularyWordsService)DeleteVocabularyWords(vocWords ivocabulary.VocabularyWords) (err error) {
	err = global.GVA_DB.Delete(&vocWords).Error
	return err
}

// DeleteVocabularyWordsByIds 批量删除VocabularyWords记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocWordsService *VocabularyWordsService)DeleteVocabularyWordsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ivocabulary.VocabularyWords{},"id in ?",ids.Ids).Error
	return err
}

// UpdateVocabularyWords 更新VocabularyWords记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocWordsService *VocabularyWordsService)UpdateVocabularyWords(vocWords ivocabulary.VocabularyWords) (err error) {
	err = global.GVA_DB.Save(&vocWords).Error
	return err
}

// GetVocabularyWords 根据id获取VocabularyWords记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocWordsService *VocabularyWordsService)GetVocabularyWords(id uint) (err error, vocWords ivocabulary.VocabularyWords) {
	err = global.GVA_DB.Where("id = ?", id).First(&vocWords).Error
	return
}

// GetVocabularyWordsInfoList 分页获取VocabularyWords记录
// Author [piexlmax](https://github.com/piexlmax)
func (vocWordsService *VocabularyWordsService)GetVocabularyWordsInfoList(info ivocabularyReq.VocabularyWordsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ivocabulary.VocabularyWords{})
    var vocWordss []ivocabulary.VocabularyWords
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&vocWordss).Error
	return err, vocWordss, total
}

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
)

type WordEgsService struct {
}

// CreateWordEgs 创建WordEgs记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordEgsService *WordEgsService) CreateWordEgs(wordEgs ivocabulary.WordEgs) (err error) {
	err = global.GVA_DB.Create(&wordEgs).Error
	return err
}

// DeleteWordEgs 删除WordEgs记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordEgsService *WordEgsService)DeleteWordEgs(wordEgs ivocabulary.WordEgs) (err error) {
	err = global.GVA_DB.Delete(&wordEgs).Error
	return err
}

// DeleteWordEgsByIds 批量删除WordEgs记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordEgsService *WordEgsService)DeleteWordEgsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ivocabulary.WordEgs{},"id in ?",ids.Ids).Error
	return err
}

// UpdateWordEgs 更新WordEgs记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordEgsService *WordEgsService)UpdateWordEgs(wordEgs ivocabulary.WordEgs) (err error) {
	err = global.GVA_DB.Save(&wordEgs).Error
	return err
}

// GetWordEgs 根据id获取WordEgs记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordEgsService *WordEgsService)GetWordEgs(id uint) (err error, wordEgs ivocabulary.WordEgs) {
	err = global.GVA_DB.Where("id = ?", id).First(&wordEgs).Error
	return
}

// GetWordEgsInfoList 分页获取WordEgs记录
// Author [piexlmax](https://github.com/piexlmax)
func (wordEgsService *WordEgsService)GetWordEgsInfoList(info ivocabularyReq.WordEgsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ivocabulary.WordEgs{})
    var wordEgss []ivocabulary.WordEgs
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&wordEgss).Error
	return err, wordEgss, total
}

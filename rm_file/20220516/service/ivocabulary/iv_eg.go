package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
)

type EgInfoService struct {
}

// CreateEgInfo 创建EgInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (egsService *EgInfoService) CreateEgInfo(egs ivocabulary.EgInfo) (err error) {
	err = global.GVA_DB.Create(&egs).Error
	return err
}

// DeleteEgInfo 删除EgInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (egsService *EgInfoService)DeleteEgInfo(egs ivocabulary.EgInfo) (err error) {
	err = global.GVA_DB.Delete(&egs).Error
	return err
}

// DeleteEgInfoByIds 批量删除EgInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (egsService *EgInfoService)DeleteEgInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ivocabulary.EgInfo{},"id in ?",ids.Ids).Error
	return err
}

// UpdateEgInfo 更新EgInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (egsService *EgInfoService)UpdateEgInfo(egs ivocabulary.EgInfo) (err error) {
	err = global.GVA_DB.Save(&egs).Error
	return err
}

// GetEgInfo 根据id获取EgInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (egsService *EgInfoService)GetEgInfo(id uint) (err error, egs ivocabulary.EgInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&egs).Error
	return
}

// GetEgInfoInfoList 分页获取EgInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (egsService *EgInfoService)GetEgInfoInfoList(info ivocabularyReq.EgInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ivocabulary.EgInfo{})
    var egss []ivocabulary.EgInfo
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&egss).Error
	return err, egss, total
}

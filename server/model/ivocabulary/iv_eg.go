// 自动生成模板EgInfo
package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// EgInfo 结构体
// 如果含有time.Time 请自行import time包
type EgInfo struct {
      global.GVA_MODEL
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:例句UUID;size:191;"`
      ContentEn  string `json:"contentEn" form:"contentEn" gorm:"column:content_en;comment:例句英文;size:512;"`
      ContentZh  string `json:"contentZh" form:"contentZh" gorm:"column:content_zh;comment:例句中文;size:512;"`
}


// TableName EgInfo 表名
func (EgInfo) TableName() string {
  return "iv_egs"
}


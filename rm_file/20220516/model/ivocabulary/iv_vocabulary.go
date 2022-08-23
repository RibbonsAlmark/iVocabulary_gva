// 自动生成模板VocabularyInfo
package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VocabularyInfo 结构体
// 如果含有time.Time 请自行import time包
type VocabularyInfo struct {
      global.GVA_MODEL
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:单词本UUID;size:191;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:单词本名称;size:191;"`
}


// TableName VocabularyInfo 表名
func (VocabularyInfo) TableName() string {
  return "iv_vocabularies"
}


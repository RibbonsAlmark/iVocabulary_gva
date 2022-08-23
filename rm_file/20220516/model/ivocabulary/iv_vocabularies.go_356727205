// 自动生成模板IvVocabularies
package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// IvVocabularies 结构体
// 如果含有time.Time 请自行import time包
type IvVocabularies struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:单词本名称;size:191;"`
      Uuid  *int `json:"uuid" form:"uuid" gorm:"column:uuid;comment:单词本UUID;size:64;"`
}


// TableName IvVocabularies 表名
func (IvVocabularies) TableName() string {
  return "iv_vocabularies"
}


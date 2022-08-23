// 自动生成模板WordEgs
package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WordEgs 结构体
// 如果含有time.Time 请自行import time包
type WordEgs struct {
      global.GVA_MODEL
      Uuid  *int `json:"uuid" form:"uuid" gorm:"column:uuid;comment:单词到例句的映射关系的UUID;size:64;"`
      WordUuid  *int `json:"wordUuid" form:"wordUuid" gorm:"column:word_uuid;comment:单词UUID;size:64;"`
      EgUuid  *int `json:"egUuid" form:"egUuid" gorm:"column:eg_uuid;comment:例句UUID;size:64;"`
}


// TableName WordEgs 表名
func (WordEgs) TableName() string {
  return "iv_word_egs"
}


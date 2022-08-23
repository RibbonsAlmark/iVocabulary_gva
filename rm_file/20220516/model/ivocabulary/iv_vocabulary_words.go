// 自动生成模板VocabularyWords
package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// VocabularyWords 结构体
// 如果含有time.Time 请自行import time包
type VocabularyWords struct {
      global.GVA_MODEL
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:映射关系UUID;size:191;"`
      VocabularyUuid  string `json:"vocabularyUuid" form:"vocabularyUuid" gorm:"column:vocabulary_uuid;comment:单词本UUID;size:191;"`
      WordUuid  string `json:"wordUuid" form:"wordUuid" gorm:"column:word_uuid;comment:单词UUID;size:191;"`
}


// TableName VocabularyWords 表名
func (VocabularyWords) TableName() string {
  return "iv_vocabulary_words"
}


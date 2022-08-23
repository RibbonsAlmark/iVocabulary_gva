// 自动生成模板WordInfo
package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// WordInfo 结构体
// 如果含有time.Time 请自行import time包
type WordInfo struct {
      global.GVA_MODEL
      Uuid  *int `json:"uuid" form:"uuid" gorm:"column:uuid;comment:单词UUID;size:64;"`
      Word  string `json:"word" form:"word" gorm:"column:word;comment:英文单词;size:128;"`
      PhoneticEn  string `json:"phoneticEn" form:"phoneticEn" gorm:"column:phonetic_en;comment:音标（英音）;size:191;"`
      PhoneticUs  string `json:"phoneticUs" form:"phoneticUs" gorm:"column:phonetic_us;comment:音标（美音）;size:191;"`
      Meanings  string `json:"meanings" form:"meanings" gorm:"column:meanings;comment:词义;size:512;"`
      ExamInfo  string `json:"examInfo" form:"examInfo" gorm:"column:exam_info;comment:在哪些考试中使用;size:191;"`
      Mark  *bool `json:"mark" form:"mark" gorm:"column:mark;comment:标记;size:1;"`
      MarkCount  *int `json:"markCount" form:"markCount" gorm:"column:mark_count;comment:被标记次数;size:128;"`
      Notes  string `json:"notes" form:"notes" gorm:"column:notes;comment:笔记;size:512;"`
      AudioFpAm  string `json:"audioFpAm" form:"audioFpAm" gorm:"column:audio_fp_am;comment:读音文件路径（英音）;size:191;"`
      AudioFpEn  string `json:"audioFpEn" form:"audioFpEn" gorm:"column:audio_fp_en;comment:读音文件路径（美音）;size:191;"`
      AudioFpTts  string `json:"audioFpTts" form:"audioFpTts" gorm:"column:audio_fp_tts;comment:读音文件路径（tts）;size:191;"`
}


// TableName WordInfo 表名
func (WordInfo) TableName() string {
  return "iv_words"
}


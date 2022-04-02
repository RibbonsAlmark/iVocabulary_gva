package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type VocabularyInfoSearch struct{
    ivocabulary.VocabularyInfo
    request.PageInfo
}

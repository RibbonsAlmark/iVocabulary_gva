package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type WordInfoSearch struct{
    ivocabulary.WordInfo
    request.PageInfo
}

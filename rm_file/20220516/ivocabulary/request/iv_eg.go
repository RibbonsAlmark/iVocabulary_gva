package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type EgInfoSearch struct{
    ivocabulary.EgInfo
    request.PageInfo
}

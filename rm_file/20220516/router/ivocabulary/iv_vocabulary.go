package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VocabularyInfoRouter struct {
}

// InitVocabularyInfoRouter 初始化 VocabularyInfo 路由信息
func (s *VocabularyInfoRouter) InitVocabularyInfoRouter(Router *gin.RouterGroup) {
	vocRouter := Router.Group("voc").Use(middleware.OperationRecord())
	vocRouterWithoutRecord := Router.Group("voc")
	var vocApi = v1.ApiGroupApp.IvocabularyApiGroup.VocabularyInfoApi
	{
		vocRouter.POST("createVocabularyInfo", vocApi.CreateVocabularyInfo)   // 新建VocabularyInfo
		vocRouter.DELETE("deleteVocabularyInfo", vocApi.DeleteVocabularyInfo) // 删除VocabularyInfo
		vocRouter.DELETE("deleteVocabularyInfoByIds", vocApi.DeleteVocabularyInfoByIds) // 批量删除VocabularyInfo
		vocRouter.PUT("updateVocabularyInfo", vocApi.UpdateVocabularyInfo)    // 更新VocabularyInfo
	}
	{
		vocRouterWithoutRecord.GET("findVocabularyInfo", vocApi.FindVocabularyInfo)        // 根据ID获取VocabularyInfo
		vocRouterWithoutRecord.GET("getVocabularyInfoList", vocApi.GetVocabularyInfoList)  // 获取VocabularyInfo列表
	}
}

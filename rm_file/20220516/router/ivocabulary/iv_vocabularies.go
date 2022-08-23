package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type IvVocabulariesRouter struct {
}

// InitIvVocabulariesRouter 初始化 IvVocabularies 路由信息
func (s *IvVocabulariesRouter) InitIvVocabulariesRouter(Router *gin.RouterGroup) {
	ivVocabulariesRouter := Router.Group("ivVocabularies").Use(middleware.OperationRecord())
	ivVocabulariesRouterWithoutRecord := Router.Group("ivVocabularies")
	var ivVocabulariesApi = v1.ApiGroupApp.IvocabularyApiGroup.IvVocabulariesApi
	{
		ivVocabulariesRouter.POST("createIvVocabularies", ivVocabulariesApi.CreateIvVocabularies)   // 新建IvVocabularies
		ivVocabulariesRouter.DELETE("deleteIvVocabularies", ivVocabulariesApi.DeleteIvVocabularies) // 删除IvVocabularies
		ivVocabulariesRouter.DELETE("deleteIvVocabulariesByIds", ivVocabulariesApi.DeleteIvVocabulariesByIds) // 批量删除IvVocabularies
		ivVocabulariesRouter.PUT("updateIvVocabularies", ivVocabulariesApi.UpdateIvVocabularies)    // 更新IvVocabularies
	}
	{
		ivVocabulariesRouterWithoutRecord.GET("findIvVocabularies", ivVocabulariesApi.FindIvVocabularies)        // 根据ID获取IvVocabularies
		ivVocabulariesRouterWithoutRecord.GET("getIvVocabulariesList", ivVocabulariesApi.GetIvVocabulariesList)  // 获取IvVocabularies列表
	}
}

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type VocabularyWordsRouter struct {
}

// InitVocabularyWordsRouter 初始化 VocabularyWords 路由信息
func (s *VocabularyWordsRouter) InitVocabularyWordsRouter(Router *gin.RouterGroup) {
	vocWordsRouter := Router.Group("vocWords").Use(middleware.OperationRecord())
	vocWordsRouterWithoutRecord := Router.Group("vocWords")
	var vocWordsApi = v1.ApiGroupApp.IvocabularyApiGroup.VocabularyWordsApi
	{
		vocWordsRouter.POST("createVocabularyWords", vocWordsApi.CreateVocabularyWords)   // 新建VocabularyWords
		vocWordsRouter.DELETE("deleteVocabularyWords", vocWordsApi.DeleteVocabularyWords) // 删除VocabularyWords
		vocWordsRouter.DELETE("deleteVocabularyWordsByIds", vocWordsApi.DeleteVocabularyWordsByIds) // 批量删除VocabularyWords
		vocWordsRouter.PUT("updateVocabularyWords", vocWordsApi.UpdateVocabularyWords)    // 更新VocabularyWords
	}
	{
		vocWordsRouterWithoutRecord.GET("findVocabularyWords", vocWordsApi.FindVocabularyWords)        // 根据ID获取VocabularyWords
		vocWordsRouterWithoutRecord.GET("getVocabularyWordsList", vocWordsApi.GetVocabularyWordsList)  // 获取VocabularyWords列表
	}
}

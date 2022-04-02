package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WordInfoRouter struct {
}

// InitWordInfoRouter 初始化 WordInfo 路由信息
func (s *WordInfoRouter) InitWordInfoRouter(Router *gin.RouterGroup) {
	wordsRouter := Router.Group("words").Use(middleware.OperationRecord())
	wordsRouterWithoutRecord := Router.Group("words")
	var wordsApi = v1.ApiGroupApp.IvocabularyApiGroup.WordInfoApi
	{
		wordsRouter.POST("createWordInfo", wordsApi.CreateWordInfo)   // 新建WordInfo
		wordsRouter.DELETE("deleteWordInfo", wordsApi.DeleteWordInfo) // 删除WordInfo
		wordsRouter.DELETE("deleteWordInfoByIds", wordsApi.DeleteWordInfoByIds) // 批量删除WordInfo
		wordsRouter.PUT("updateWordInfo", wordsApi.UpdateWordInfo)    // 更新WordInfo
	}
	{
		wordsRouterWithoutRecord.GET("findWordInfo", wordsApi.FindWordInfo)        // 根据ID获取WordInfo
		wordsRouterWithoutRecord.GET("getWordInfoList", wordsApi.GetWordInfoList)  // 获取WordInfo列表
	}
}

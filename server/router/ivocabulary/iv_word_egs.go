package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WordEgsRouter struct {
}

// InitWordEgsRouter 初始化 WordEgs 路由信息
func (s *WordEgsRouter) InitWordEgsRouter(Router *gin.RouterGroup) {
	wordEgsRouter := Router.Group("wordEgs").Use(middleware.OperationRecord())
	wordEgsRouterWithoutRecord := Router.Group("wordEgs")
	var wordEgsApi = v1.ApiGroupApp.IvocabularyApiGroup.WordEgsApi
	{
		wordEgsRouter.POST("createWordEgs", wordEgsApi.CreateWordEgs)   // 新建WordEgs
		wordEgsRouter.DELETE("deleteWordEgs", wordEgsApi.DeleteWordEgs) // 删除WordEgs
		wordEgsRouter.DELETE("deleteWordEgsByIds", wordEgsApi.DeleteWordEgsByIds) // 批量删除WordEgs
		wordEgsRouter.PUT("updateWordEgs", wordEgsApi.UpdateWordEgs)    // 更新WordEgs
	}
	{
		wordEgsRouterWithoutRecord.GET("findWordEgs", wordEgsApi.FindWordEgs)        // 根据ID获取WordEgs
		wordEgsRouterWithoutRecord.GET("getWordEgsList", wordEgsApi.GetWordEgsList)  // 获取WordEgs列表
	}
}

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EgInfoRouter struct {
}

// InitEgInfoRouter 初始化 EgInfo 路由信息
func (s *EgInfoRouter) InitEgInfoRouter(Router *gin.RouterGroup) {
	egsRouter := Router.Group("egs").Use(middleware.OperationRecord())
	egsRouterWithoutRecord := Router.Group("egs")
	var egsApi = v1.ApiGroupApp.IvocabularyApiGroup.EgInfoApi
	{
		egsRouter.POST("createEgInfo", egsApi.CreateEgInfo)   // 新建EgInfo
		egsRouter.DELETE("deleteEgInfo", egsApi.DeleteEgInfo) // 删除EgInfo
		egsRouter.DELETE("deleteEgInfoByIds", egsApi.DeleteEgInfoByIds) // 批量删除EgInfo
		egsRouter.PUT("updateEgInfo", egsApi.UpdateEgInfo)    // 更新EgInfo
	}
	{
		egsRouterWithoutRecord.GET("findEgInfo", egsApi.FindEgInfo)        // 根据ID获取EgInfo
		egsRouterWithoutRecord.GET("getEgInfoList", egsApi.GetEgInfoList)  // 获取EgInfo列表
	}
}

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type EgInfoApi struct {
}

var egsService = service.ServiceGroupApp.IvocabularyServiceGroup.EgInfoService


// CreateEgInfo 创建EgInfo
// @Tags EgInfo
// @Summary 创建EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.EgInfo true "创建EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /egs/createEgInfo [post]
func (egsApi *EgInfoApi) CreateEgInfo(c *gin.Context) {
	var egs ivocabulary.EgInfo
	_ = c.ShouldBindJSON(&egs)
	if err := egsService.CreateEgInfo(egs); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteEgInfo 删除EgInfo
// @Tags EgInfo
// @Summary 删除EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.EgInfo true "删除EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /egs/deleteEgInfo [delete]
func (egsApi *EgInfoApi) DeleteEgInfo(c *gin.Context) {
	var egs ivocabulary.EgInfo
	_ = c.ShouldBindJSON(&egs)
	if err := egsService.DeleteEgInfo(egs); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteEgInfoByIds 批量删除EgInfo
// @Tags EgInfo
// @Summary 批量删除EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /egs/deleteEgInfoByIds [delete]
func (egsApi *EgInfoApi) DeleteEgInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := egsService.DeleteEgInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateEgInfo 更新EgInfo
// @Tags EgInfo
// @Summary 更新EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.EgInfo true "更新EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /egs/updateEgInfo [put]
func (egsApi *EgInfoApi) UpdateEgInfo(c *gin.Context) {
	var egs ivocabulary.EgInfo
	_ = c.ShouldBindJSON(&egs)
	if err := egsService.UpdateEgInfo(egs); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindEgInfo 用id查询EgInfo
// @Tags EgInfo
// @Summary 用id查询EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabulary.EgInfo true "用id查询EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /egs/findEgInfo [get]
func (egsApi *EgInfoApi) FindEgInfo(c *gin.Context) {
	var egs ivocabulary.EgInfo
	_ = c.ShouldBindQuery(&egs)
	if err, reegs := egsService.GetEgInfo(egs.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reegs": reegs}, c)
	}
}

// GetEgInfoList 分页获取EgInfo列表
// @Tags EgInfo
// @Summary 分页获取EgInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabularyReq.EgInfoSearch true "分页获取EgInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /egs/getEgInfoList [get]
func (egsApi *EgInfoApi) GetEgInfoList(c *gin.Context) {
	var pageInfo ivocabularyReq.EgInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := egsService.GetEgInfoInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

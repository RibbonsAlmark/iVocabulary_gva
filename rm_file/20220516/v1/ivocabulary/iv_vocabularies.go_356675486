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

type IvVocabulariesApi struct {
}

var ivVocabulariesService = service.ServiceGroupApp.IvocabularyServiceGroup.IvVocabulariesService


// CreateIvVocabularies 创建IvVocabularies
// @Tags IvVocabularies
// @Summary 创建IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.IvVocabularies true "创建IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ivVocabularies/createIvVocabularies [post]
func (ivVocabulariesApi *IvVocabulariesApi) CreateIvVocabularies(c *gin.Context) {
	var ivVocabularies ivocabulary.IvVocabularies
	_ = c.ShouldBindJSON(&ivVocabularies)
	if err := ivVocabulariesService.CreateIvVocabularies(ivVocabularies); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteIvVocabularies 删除IvVocabularies
// @Tags IvVocabularies
// @Summary 删除IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.IvVocabularies true "删除IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ivVocabularies/deleteIvVocabularies [delete]
func (ivVocabulariesApi *IvVocabulariesApi) DeleteIvVocabularies(c *gin.Context) {
	var ivVocabularies ivocabulary.IvVocabularies
	_ = c.ShouldBindJSON(&ivVocabularies)
	if err := ivVocabulariesService.DeleteIvVocabularies(ivVocabularies); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteIvVocabulariesByIds 批量删除IvVocabularies
// @Tags IvVocabularies
// @Summary 批量删除IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ivVocabularies/deleteIvVocabulariesByIds [delete]
func (ivVocabulariesApi *IvVocabulariesApi) DeleteIvVocabulariesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := ivVocabulariesService.DeleteIvVocabulariesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateIvVocabularies 更新IvVocabularies
// @Tags IvVocabularies
// @Summary 更新IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.IvVocabularies true "更新IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ivVocabularies/updateIvVocabularies [put]
func (ivVocabulariesApi *IvVocabulariesApi) UpdateIvVocabularies(c *gin.Context) {
	var ivVocabularies ivocabulary.IvVocabularies
	_ = c.ShouldBindJSON(&ivVocabularies)
	if err := ivVocabulariesService.UpdateIvVocabularies(ivVocabularies); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindIvVocabularies 用id查询IvVocabularies
// @Tags IvVocabularies
// @Summary 用id查询IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabulary.IvVocabularies true "用id查询IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ivVocabularies/findIvVocabularies [get]
func (ivVocabulariesApi *IvVocabulariesApi) FindIvVocabularies(c *gin.Context) {
	var ivVocabularies ivocabulary.IvVocabularies
	_ = c.ShouldBindQuery(&ivVocabularies)
	if err, reivVocabularies := ivVocabulariesService.GetIvVocabularies(ivVocabularies.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reivVocabularies": reivVocabularies}, c)
	}
}

// GetIvVocabulariesList 分页获取IvVocabularies列表
// @Tags IvVocabularies
// @Summary 分页获取IvVocabularies列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabularyReq.IvVocabulariesSearch true "分页获取IvVocabularies列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ivVocabularies/getIvVocabulariesList [get]
func (ivVocabulariesApi *IvVocabulariesApi) GetIvVocabulariesList(c *gin.Context) {
	var pageInfo ivocabularyReq.IvVocabulariesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := ivVocabulariesService.GetIvVocabulariesInfoList(pageInfo); err != nil {
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

package ivocabulary

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary"
	ivocabularyReq "github.com/flipped-aurora/gin-vue-admin/server/model/ivocabulary/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type VocabularyInfoApi struct {
}

var vocService = service.ServiceGroupApp.IvocabularyServiceGroup.VocabularyInfoService

// CreateVocabularyInfo 创建VocabularyInfo
// @Tags VocabularyInfo
// @Summary 创建VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.VocabularyInfo true "创建VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /voc/createVocabularyInfo [post]
func (vocApi *VocabularyInfoApi) CreateVocabularyInfo(c *gin.Context) {
	var voc ivocabulary.VocabularyInfo
	_ = c.ShouldBindJSON(&voc)
	if err := vocService.CreateVocabularyInfo(voc); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVocabularyInfo 删除VocabularyInfo
// @Tags VocabularyInfo
// @Summary 删除VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.VocabularyInfo true "删除VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /voc/deleteVocabularyInfo [delete]
func (vocApi *VocabularyInfoApi) DeleteVocabularyInfo(c *gin.Context) {
	var voc ivocabulary.VocabularyInfo
	_ = c.ShouldBindJSON(&voc)
	if err := vocService.DeleteVocabularyInfo(voc); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVocabularyInfoByIds 批量删除VocabularyInfo
// @Tags VocabularyInfo
// @Summary 批量删除VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /voc/deleteVocabularyInfoByIds [delete]
func (vocApi *VocabularyInfoApi) DeleteVocabularyInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := vocService.DeleteVocabularyInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVocabularyInfo 更新VocabularyInfo
// @Tags VocabularyInfo
// @Summary 更新VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.VocabularyInfo true "更新VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /voc/updateVocabularyInfo [put]
func (vocApi *VocabularyInfoApi) UpdateVocabularyInfo(c *gin.Context) {
	var voc ivocabulary.VocabularyInfo
	_ = c.ShouldBindJSON(&voc)
	if err := vocService.UpdateVocabularyInfo(voc); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVocabularyInfo 用id查询VocabularyInfo
// @Tags VocabularyInfo
// @Summary 用id查询VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabulary.VocabularyInfo true "用id查询VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /voc/findVocabularyInfo [get]
func (vocApi *VocabularyInfoApi) FindVocabularyInfo(c *gin.Context) {
	var voc ivocabulary.VocabularyInfo
	_ = c.ShouldBindQuery(&voc)
	if err, revoc := vocService.GetVocabularyInfo(voc.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revoc": revoc}, c)
	}
}

// GetVocabularyInfoList 分页获取VocabularyInfo列表
// @Tags VocabularyInfo
// @Summary 分页获取VocabularyInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabularyReq.VocabularyInfoSearch true "分页获取VocabularyInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /voc/getVocabularyInfoList [get]
func (vocApi *VocabularyInfoApi) GetVocabularyInfoList(c *gin.Context) {
	var pageInfo ivocabularyReq.VocabularyInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := vocService.GetVocabularyInfoInfoList(pageInfo); err != nil {
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

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

type VocabularyWordsApi struct {
}

var vocWordsService = service.ServiceGroupApp.IvocabularyServiceGroup.VocabularyWordsService


// CreateVocabularyWords 创建VocabularyWords
// @Tags VocabularyWords
// @Summary 创建VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.VocabularyWords true "创建VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vocWords/createVocabularyWords [post]
func (vocWordsApi *VocabularyWordsApi) CreateVocabularyWords(c *gin.Context) {
	var vocWords ivocabulary.VocabularyWords
	_ = c.ShouldBindJSON(&vocWords)
	if err := vocWordsService.CreateVocabularyWords(vocWords); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteVocabularyWords 删除VocabularyWords
// @Tags VocabularyWords
// @Summary 删除VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.VocabularyWords true "删除VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vocWords/deleteVocabularyWords [delete]
func (vocWordsApi *VocabularyWordsApi) DeleteVocabularyWords(c *gin.Context) {
	var vocWords ivocabulary.VocabularyWords
	_ = c.ShouldBindJSON(&vocWords)
	if err := vocWordsService.DeleteVocabularyWords(vocWords); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteVocabularyWordsByIds 批量删除VocabularyWords
// @Tags VocabularyWords
// @Summary 批量删除VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /vocWords/deleteVocabularyWordsByIds [delete]
func (vocWordsApi *VocabularyWordsApi) DeleteVocabularyWordsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := vocWordsService.DeleteVocabularyWordsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateVocabularyWords 更新VocabularyWords
// @Tags VocabularyWords
// @Summary 更新VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.VocabularyWords true "更新VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vocWords/updateVocabularyWords [put]
func (vocWordsApi *VocabularyWordsApi) UpdateVocabularyWords(c *gin.Context) {
	var vocWords ivocabulary.VocabularyWords
	_ = c.ShouldBindJSON(&vocWords)
	if err := vocWordsService.UpdateVocabularyWords(vocWords); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindVocabularyWords 用id查询VocabularyWords
// @Tags VocabularyWords
// @Summary 用id查询VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabulary.VocabularyWords true "用id查询VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vocWords/findVocabularyWords [get]
func (vocWordsApi *VocabularyWordsApi) FindVocabularyWords(c *gin.Context) {
	var vocWords ivocabulary.VocabularyWords
	_ = c.ShouldBindQuery(&vocWords)
	if err, revocWords := vocWordsService.GetVocabularyWords(vocWords.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"revocWords": revocWords}, c)
	}
}

// GetVocabularyWordsList 分页获取VocabularyWords列表
// @Tags VocabularyWords
// @Summary 分页获取VocabularyWords列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabularyReq.VocabularyWordsSearch true "分页获取VocabularyWords列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vocWords/getVocabularyWordsList [get]
func (vocWordsApi *VocabularyWordsApi) GetVocabularyWordsList(c *gin.Context) {
	var pageInfo ivocabularyReq.VocabularyWordsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := vocWordsService.GetVocabularyWordsInfoList(pageInfo); err != nil {
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

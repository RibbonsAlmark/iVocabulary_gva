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

type WordInfoApi struct {
}

var wordsService = service.ServiceGroupApp.IvocabularyServiceGroup.WordInfoService


// CreateWordInfo 创建WordInfo
// @Tags WordInfo
// @Summary 创建WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.WordInfo true "创建WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /words/createWordInfo [post]
func (wordsApi *WordInfoApi) CreateWordInfo(c *gin.Context) {
	var words ivocabulary.WordInfo
	_ = c.ShouldBindJSON(&words)
	if err := wordsService.CreateWordInfo(words); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWordInfo 删除WordInfo
// @Tags WordInfo
// @Summary 删除WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.WordInfo true "删除WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /words/deleteWordInfo [delete]
func (wordsApi *WordInfoApi) DeleteWordInfo(c *gin.Context) {
	var words ivocabulary.WordInfo
	_ = c.ShouldBindJSON(&words)
	if err := wordsService.DeleteWordInfo(words); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWordInfoByIds 批量删除WordInfo
// @Tags WordInfo
// @Summary 批量删除WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /words/deleteWordInfoByIds [delete]
func (wordsApi *WordInfoApi) DeleteWordInfoByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := wordsService.DeleteWordInfoByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWordInfo 更新WordInfo
// @Tags WordInfo
// @Summary 更新WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.WordInfo true "更新WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /words/updateWordInfo [put]
func (wordsApi *WordInfoApi) UpdateWordInfo(c *gin.Context) {
	var words ivocabulary.WordInfo
	_ = c.ShouldBindJSON(&words)
	if err := wordsService.UpdateWordInfo(words); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWordInfo 用id查询WordInfo
// @Tags WordInfo
// @Summary 用id查询WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabulary.WordInfo true "用id查询WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /words/findWordInfo [get]
func (wordsApi *WordInfoApi) FindWordInfo(c *gin.Context) {
	var words ivocabulary.WordInfo
	_ = c.ShouldBindQuery(&words)
	if err, rewords := wordsService.GetWordInfo(words.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewords": rewords}, c)
	}
}

// GetWordInfoList 分页获取WordInfo列表
// @Tags WordInfo
// @Summary 分页获取WordInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabularyReq.WordInfoSearch true "分页获取WordInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /words/getWordInfoList [get]
func (wordsApi *WordInfoApi) GetWordInfoList(c *gin.Context) {
	var pageInfo ivocabularyReq.WordInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := wordsService.GetWordInfoInfoList(pageInfo); err != nil {
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

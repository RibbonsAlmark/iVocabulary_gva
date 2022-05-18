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

type WordEgsApi struct {
}

var wordEgsService = service.ServiceGroupApp.IvocabularyServiceGroup.WordEgsService


// CreateWordEgs 创建WordEgs
// @Tags WordEgs
// @Summary 创建WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.WordEgs true "创建WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wordEgs/createWordEgs [post]
func (wordEgsApi *WordEgsApi) CreateWordEgs(c *gin.Context) {
	var wordEgs ivocabulary.WordEgs
	_ = c.ShouldBindJSON(&wordEgs)
	if err := wordEgsService.CreateWordEgs(wordEgs); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWordEgs 删除WordEgs
// @Tags WordEgs
// @Summary 删除WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.WordEgs true "删除WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wordEgs/deleteWordEgs [delete]
func (wordEgsApi *WordEgsApi) DeleteWordEgs(c *gin.Context) {
	var wordEgs ivocabulary.WordEgs
	_ = c.ShouldBindJSON(&wordEgs)
	if err := wordEgsService.DeleteWordEgs(wordEgs); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWordEgsByIds 批量删除WordEgs
// @Tags WordEgs
// @Summary 批量删除WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /wordEgs/deleteWordEgsByIds [delete]
func (wordEgsApi *WordEgsApi) DeleteWordEgsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := wordEgsService.DeleteWordEgsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWordEgs 更新WordEgs
// @Tags WordEgs
// @Summary 更新WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ivocabulary.WordEgs true "更新WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wordEgs/updateWordEgs [put]
func (wordEgsApi *WordEgsApi) UpdateWordEgs(c *gin.Context) {
	var wordEgs ivocabulary.WordEgs
	_ = c.ShouldBindJSON(&wordEgs)
	if err := wordEgsService.UpdateWordEgs(wordEgs); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWordEgs 用id查询WordEgs
// @Tags WordEgs
// @Summary 用id查询WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabulary.WordEgs true "用id查询WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wordEgs/findWordEgs [get]
func (wordEgsApi *WordEgsApi) FindWordEgs(c *gin.Context) {
	var wordEgs ivocabulary.WordEgs
	_ = c.ShouldBindQuery(&wordEgs)
	if err, rewordEgs := wordEgsService.GetWordEgs(wordEgs.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewordEgs": rewordEgs}, c)
	}
}

// GetWordEgsList 分页获取WordEgs列表
// @Tags WordEgs
// @Summary 分页获取WordEgs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ivocabularyReq.WordEgsSearch true "分页获取WordEgs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wordEgs/getWordEgsList [get]
func (wordEgsApi *WordEgsApi) GetWordEgsList(c *gin.Context) {
	var pageInfo ivocabularyReq.WordEgsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := wordEgsService.GetWordEgsInfoList(pageInfo); err != nil {
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

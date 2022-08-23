import service from '@/utils/request'

// @Tags WordInfo
// @Summary 创建WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WordInfo true "创建WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /words/createWordInfo [post]
export const createWordInfo = (data) => {
  return service({
    url: '/words/createWordInfo',
    method: 'post',
    data
  })
}

// @Tags WordInfo
// @Summary 删除WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WordInfo true "删除WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /words/deleteWordInfo [delete]
export const deleteWordInfo = (data) => {
  return service({
    url: '/words/deleteWordInfo',
    method: 'delete',
    data
  })
}

// @Tags WordInfo
// @Summary 删除WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /words/deleteWordInfo [delete]
export const deleteWordInfoByIds = (data) => {
  return service({
    url: '/words/deleteWordInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags WordInfo
// @Summary 更新WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WordInfo true "更新WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /words/updateWordInfo [put]
export const updateWordInfo = (data) => {
  return service({
    url: '/words/updateWordInfo',
    method: 'put',
    data
  })
}

// @Tags WordInfo
// @Summary 用id查询WordInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WordInfo true "用id查询WordInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /words/findWordInfo [get]
export const findWordInfo = (params) => {
  return service({
    url: '/words/findWordInfo',
    method: 'get',
    params
  })
}

// @Tags WordInfo
// @Summary 分页获取WordInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WordInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /words/getWordInfoList [get]
export const getWordInfoList = (params) => {
  return service({
    url: '/words/getWordInfoList',
    method: 'get',
    params
  })
}

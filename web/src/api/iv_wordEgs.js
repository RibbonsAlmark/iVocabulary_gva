import service from '@/utils/request'

// @Tags WordEgs
// @Summary 创建WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WordEgs true "创建WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wordEgs/createWordEgs [post]
export const createWordEgs = (data) => {
  return service({
    url: '/wordEgs/createWordEgs',
    method: 'post',
    data
  })
}

// @Tags WordEgs
// @Summary 删除WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WordEgs true "删除WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wordEgs/deleteWordEgs [delete]
export const deleteWordEgs = (data) => {
  return service({
    url: '/wordEgs/deleteWordEgs',
    method: 'delete',
    data
  })
}

// @Tags WordEgs
// @Summary 删除WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /wordEgs/deleteWordEgs [delete]
export const deleteWordEgsByIds = (data) => {
  return service({
    url: '/wordEgs/deleteWordEgsByIds',
    method: 'delete',
    data
  })
}

// @Tags WordEgs
// @Summary 更新WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WordEgs true "更新WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /wordEgs/updateWordEgs [put]
export const updateWordEgs = (data) => {
  return service({
    url: '/wordEgs/updateWordEgs',
    method: 'put',
    data
  })
}

// @Tags WordEgs
// @Summary 用id查询WordEgs
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WordEgs true "用id查询WordEgs"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /wordEgs/findWordEgs [get]
export const findWordEgs = (params) => {
  return service({
    url: '/wordEgs/findWordEgs',
    method: 'get',
    params
  })
}

// @Tags WordEgs
// @Summary 分页获取WordEgs列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取WordEgs列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /wordEgs/getWordEgsList [get]
export const getWordEgsList = (params) => {
  return service({
    url: '/wordEgs/getWordEgsList',
    method: 'get',
    params
  })
}

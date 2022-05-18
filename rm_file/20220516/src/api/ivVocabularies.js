import service from '@/utils/request'

// @Tags IvVocabularies
// @Summary 创建IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IvVocabularies true "创建IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ivVocabularies/createIvVocabularies [post]
export const createIvVocabularies = (data) => {
  return service({
    url: '/ivVocabularies/createIvVocabularies',
    method: 'post',
    data
  })
}

// @Tags IvVocabularies
// @Summary 删除IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IvVocabularies true "删除IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ivVocabularies/deleteIvVocabularies [delete]
export const deleteIvVocabularies = (data) => {
  return service({
    url: '/ivVocabularies/deleteIvVocabularies',
    method: 'delete',
    data
  })
}

// @Tags IvVocabularies
// @Summary 删除IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ivVocabularies/deleteIvVocabularies [delete]
export const deleteIvVocabulariesByIds = (data) => {
  return service({
    url: '/ivVocabularies/deleteIvVocabulariesByIds',
    method: 'delete',
    data
  })
}

// @Tags IvVocabularies
// @Summary 更新IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.IvVocabularies true "更新IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ivVocabularies/updateIvVocabularies [put]
export const updateIvVocabularies = (data) => {
  return service({
    url: '/ivVocabularies/updateIvVocabularies',
    method: 'put',
    data
  })
}

// @Tags IvVocabularies
// @Summary 用id查询IvVocabularies
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.IvVocabularies true "用id查询IvVocabularies"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ivVocabularies/findIvVocabularies [get]
export const findIvVocabularies = (params) => {
  return service({
    url: '/ivVocabularies/findIvVocabularies',
    method: 'get',
    params
  })
}

// @Tags IvVocabularies
// @Summary 分页获取IvVocabularies列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取IvVocabularies列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ivVocabularies/getIvVocabulariesList [get]
export const getIvVocabulariesList = (params) => {
  return service({
    url: '/ivVocabularies/getIvVocabulariesList',
    method: 'get',
    params
  })
}

import service from '@/utils/request'

// @Tags VocabularyWords
// @Summary 创建VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VocabularyWords true "创建VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vocWords/createVocabularyWords [post]
export const createVocabularyWords = (data) => {
  return service({
    url: '/vocWords/createVocabularyWords',
    method: 'post',
    data
  })
}

// @Tags VocabularyWords
// @Summary 删除VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VocabularyWords true "删除VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vocWords/deleteVocabularyWords [delete]
export const deleteVocabularyWords = (data) => {
  return service({
    url: '/vocWords/deleteVocabularyWords',
    method: 'delete',
    data
  })
}

// @Tags VocabularyWords
// @Summary 删除VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /vocWords/deleteVocabularyWords [delete]
export const deleteVocabularyWordsByIds = (data) => {
  return service({
    url: '/vocWords/deleteVocabularyWordsByIds',
    method: 'delete',
    data
  })
}

// @Tags VocabularyWords
// @Summary 更新VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VocabularyWords true "更新VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /vocWords/updateVocabularyWords [put]
export const updateVocabularyWords = (data) => {
  return service({
    url: '/vocWords/updateVocabularyWords',
    method: 'put',
    data
  })
}

// @Tags VocabularyWords
// @Summary 用id查询VocabularyWords
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VocabularyWords true "用id查询VocabularyWords"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /vocWords/findVocabularyWords [get]
export const findVocabularyWords = (params) => {
  return service({
    url: '/vocWords/findVocabularyWords',
    method: 'get',
    params
  })
}

// @Tags VocabularyWords
// @Summary 分页获取VocabularyWords列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VocabularyWords列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /vocWords/getVocabularyWordsList [get]
export const getVocabularyWordsList = (params) => {
  return service({
    url: '/vocWords/getVocabularyWordsList',
    method: 'get',
    params
  })
}

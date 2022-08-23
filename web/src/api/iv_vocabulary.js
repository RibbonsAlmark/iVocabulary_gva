import service from '@/utils/request'

// @Tags VocabularyInfo
// @Summary 创建VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VocabularyInfo true "创建VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /voc/createVocabularyInfo [post]
export const createVocabularyInfo = (data) => {
  return service({
    url: '/voc/createVocabularyInfo',
    method: 'post',
    data
  })
}

// @Tags VocabularyInfo
// @Summary 删除VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VocabularyInfo true "删除VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /voc/deleteVocabularyInfo [delete]
export const deleteVocabularyInfo = (data) => {
  return service({
    url: '/voc/deleteVocabularyInfo',
    method: 'delete',
    data
  })
}

// @Tags VocabularyInfo
// @Summary 删除VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /voc/deleteVocabularyInfo [delete]
export const deleteVocabularyInfoByIds = (data) => {
  return service({
    url: '/voc/deleteVocabularyInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags VocabularyInfo
// @Summary 更新VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.VocabularyInfo true "更新VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /voc/updateVocabularyInfo [put]
export const updateVocabularyInfo = (data) => {
  return service({
    url: '/voc/updateVocabularyInfo',
    method: 'put',
    data
  })
}

// @Tags VocabularyInfo
// @Summary 用id查询VocabularyInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.VocabularyInfo true "用id查询VocabularyInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /voc/findVocabularyInfo [get]
export const findVocabularyInfo = (params) => {
  return service({
    url: '/voc/findVocabularyInfo',
    method: 'get',
    params
  })
}

// @Tags VocabularyInfo
// @Summary 分页获取VocabularyInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取VocabularyInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /voc/getVocabularyInfoList [get]
export const getVocabularyInfoList = (params) => {
  return service({
    url: '/voc/getVocabularyInfoList',
    method: 'get',
    params
  })
}

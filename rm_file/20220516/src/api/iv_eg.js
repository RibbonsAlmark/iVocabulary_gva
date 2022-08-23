import service from '@/utils/request'

// @Tags EgInfo
// @Summary 创建EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EgInfo true "创建EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /egs/createEgInfo [post]
export const createEgInfo = (data) => {
  return service({
    url: '/egs/createEgInfo',
    method: 'post',
    data
  })
}

// @Tags EgInfo
// @Summary 删除EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EgInfo true "删除EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /egs/deleteEgInfo [delete]
export const deleteEgInfo = (data) => {
  return service({
    url: '/egs/deleteEgInfo',
    method: 'delete',
    data
  })
}

// @Tags EgInfo
// @Summary 删除EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /egs/deleteEgInfo [delete]
export const deleteEgInfoByIds = (data) => {
  return service({
    url: '/egs/deleteEgInfoByIds',
    method: 'delete',
    data
  })
}

// @Tags EgInfo
// @Summary 更新EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EgInfo true "更新EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /egs/updateEgInfo [put]
export const updateEgInfo = (data) => {
  return service({
    url: '/egs/updateEgInfo',
    method: 'put',
    data
  })
}

// @Tags EgInfo
// @Summary 用id查询EgInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EgInfo true "用id查询EgInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /egs/findEgInfo [get]
export const findEgInfo = (params) => {
  return service({
    url: '/egs/findEgInfo',
    method: 'get',
    params
  })
}

// @Tags EgInfo
// @Summary 分页获取EgInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取EgInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /egs/getEgInfoList [get]
export const getEgInfoList = (params) => {
  return service({
    url: '/egs/getEgInfoList',
    method: 'get',
    params
  })
}

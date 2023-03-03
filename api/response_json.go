package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// {
//   "code": 2000,
//   "msg": "tip info",
//   "data": data,
// }

// ResponseJson 返回结构体
type ResponseJson struct {
	Status int    `json:"-"`              // 响应状态码, 最后通过 http status code 展示, 所以不做解析
	Code   int    `json:"code,omitempty"` // 内部错误码
	Msg    string `json:"msg,omitempty"`  // 提示信息
	Data   any    `json:"data,omitempty"` // 数据
}

// isEmpty 判断该结构体是否为空
//func (resp *ResponseJson) isEmpty() bool {
//	return reflect.DeepEqual(*resp, ResponseJson{})
//}

// buildStatus 如果状态为空,则设置默认值
func (resp *ResponseJson) httpResponse(ctx *gin.Context, defaultStatus int) {

	// 设置状态码默认值
	if resp.Status == 0 {
		resp.Status = defaultStatus
	}

	// 如果为空则直接返回状态码
	//if resp.isEmpty() {
	//	ctx.AbortWithStatus(resp.Status)
	//} else {
	//	ctx.AbortWithStatusJSON(resp.Status, resp)
	//}

	// 由于改成了结构体方法, 所以结构体不能为空
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func (resp *ResponseJson) Success(ctx *gin.Context) {
	resp.httpResponse(ctx, http.StatusOK)
}

func (resp *ResponseJson) Failed(ctx *gin.Context) {
	resp.httpResponse(ctx, http.StatusBadRequest)
}

func (resp *ResponseJson) ServerFail(ctx *gin.Context) {
	resp.httpResponse(ctx, http.StatusInternalServerError)
}

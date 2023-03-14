package global

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ResponseJson 返回结构体
type ResponseJson struct {
	Status int  `json:"-"` // 响应状态码, 最后通过 http status code 展示, 所以不做解析
	Head   Head `json:"head"`
	Data   any  `json:"data,omitempty"` // 数据
}

type Head struct {
	Code int    `json:"code"`          // 内部错误码
	Msg  string `json:"msg,omitempty"` // 提示信息
}

// buildStatus 如果状态为空,则设置默认值
func (resp *ResponseJson) httpResponse(ctx *gin.Context, defaultStatus int) {

	// 如果结构体中状态码未被设置, 则使用默认状态码, 对 http status code 进行修改
	if resp.Status == 0 {
		resp.Status = defaultStatus
	}

	// 由于改成了结构体方法, 所以结构体不能为空
	ctx.AbortWithStatusJSON(resp.Status, resp)
}

func (resp *ResponseJson) Success(ctx *gin.Context) {
	resp.httpResponse(ctx, http.StatusOK)
}

func (resp *ResponseJson) Failed(ctx *gin.Context) {
	Logger.Error(resp.Head.Msg)
	resp.httpResponse(ctx, http.StatusBadRequest)
}

func (resp *ResponseJson) ServerFail(ctx *gin.Context) {
	resp.httpResponse(ctx, http.StatusInternalServerError)
}

package mgin

import (
	"github.com/gin-gonic/gin"
)

type Handler struct{}

func (t *Handler) Get(ctx *gin.Context) {
	if t.Action(ctx) {
		return
	}
	t.Return501(ctx)
}

func (t *Handler) Post(ctx *gin.Context) {
	if t.Action(ctx) {
		return
	}
	t.Return501(ctx)
}

func (t *Handler) Put(ctx *gin.Context) {
	if t.Action(ctx) {
		return
	}
	t.Return501(ctx)
}

func (t *Handler) Delete(ctx *gin.Context) {
	if t.Action(ctx) {
		return
	}
	t.Return501(ctx)
}

// 需要复写此方法
func (t *Handler) Action(ctx *gin.Context) bool {
	// 检查是否有action参数，如果有则执行对应动作
	if actionType := ctx.Query("action"); actionType != "" {
		action.ExecuteAction(ctx)
		return true
	}
	return false
}

// 返回400状态码, 表示请求错误
func (t *Handler) Return400(ctx *gin.Context, err string) {
	ctx.JSON(400, Response[any]{
		Code: 400,
		Err:  err,
	})
}

// 返回401状态码, 表示未授权
func (t *Handler) Return401(ctx *gin.Context, err string) {
	ctx.JSON(401, Response[any]{
		Code: 401,
		Err:  err,
	})
}

// 返回403状态码, 表示禁止访问
func (t *Handler) Return403(ctx *gin.Context, err string) {
	ctx.JSON(403, Response[any]{
		Code: 403,
		Err:  err,
	})
}

// 返回404状态码, 表示页面找不到
func (t *Handler) Return404(ctx *gin.Context, err string) {
	ctx.JSON(404, Response[any]{
		Code: 404,
		Err:  err,
	})
}

// 返回500状态码, 表示服务器错误
func (t *Handler) Return500(ctx *gin.Context, err string) {
	ctx.JSON(500, Response[any]{
		Code: 500,
		Err:  err,
	})
}

// 返回501状态码, 表示未实现
func (t *Handler) Return501(ctx *gin.Context) {
	ctx.JSON(200, Response[any]{
		Code: 501,
		Err:  "not implemented",
		Data: gin.H{
			"method": ctx.Request.Method,
			"params": ctx.Params,
		},
	})
}

// 返回503状态码, 表示服务不可用
func (t *Handler) Return503(ctx *gin.Context, err string) {
	ctx.JSON(503, Response[any]{
		Code: 503,
		Err:  err,
	})
}

package mgin

import "github.com/gin-gonic/gin"

// RESTful 接口定义了RESTful风格的动作处理器
type RESTful interface {
	Get(ctx *gin.Context, args map[string]string) (interface{}, error)
	Post(ctx *gin.Context, args map[string]string) (interface{}, error)
	Put(ctx *gin.Context, args map[string]string) (interface{}, error)
	Delete(ctx *gin.Context, args map[string]string) (interface{}, error)
}

// 先写一个简单的接口实现, 以便直接继承
type RESTfulTs struct{}

// template 模板方法, 用于给子类实现
func (t *RESTfulTs) template(ctx *gin.Context, args map[string]string) (interface{}, error) {
	return Response[any]{
		Code: 501,
		Msg:  "方法未实现",
		Data: gin.H{
			"method": ctx.Request.Method,
			"args":   args,
		},
	}, nil
}

func (t *RESTfulTs) Get(ctx *gin.Context, args map[string]string) (interface{}, error) {
	return t.template(ctx, args)
}

func (t *RESTfulTs) Post(ctx *gin.Context, args map[string]string) (interface{}, error) {
	return t.template(ctx, args)
}

func (t *RESTfulTs) Put(ctx *gin.Context, args map[string]string) (interface{}, error) {
	return t.template(ctx, args)
}

func (t *RESTfulTs) Delete(ctx *gin.Context, args map[string]string) (interface{}, error) {
	return t.template(ctx, args)
}

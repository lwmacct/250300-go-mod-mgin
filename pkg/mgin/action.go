package mgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var action = new(Action)

// Action 结构体用于管理不同类型的动作
type Action struct {
	Handlers    map[string]ActionHandler // 存储不同类型的动作处理函数
	RestActions map[string]RESTful       // 存储RESTful风格的动作处理器
}

// ActionHandler 定义动作处理函数的类型
type ActionHandler func(ctx *gin.Context, args map[string]string) (interface{}, error)

// RegisterHandler 注册一个新的动作处理函数
func (t *Action) RegisterHandler(actionType string, handler ActionHandler) {
	t.Handlers[actionType] = handler
}

// RegisterRESTfulAction 注册一个RESTful风格的动作处理器
func (t *Action) RegisterRESTfulAction(actionType string, action RESTful) {
	t.RestActions[actionType] = action
}

// ExecuteAction 执行指定类型的动作，支持RESTful方法
func (t *Action) ExecuteAction(ctx *gin.Context) {
	// 获取动作类型
	actionType := ctx.Query("action")
	if actionType == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "action parameter is required"})
		return
	}
	// 收集所有查询参数作为动作参数
	args := make(map[string]string)
	for k, v := range ctx.Request.URL.Query() {
		if k != "action" && len(v) > 0 {
			args[k] = v[0]
		}
	}

	// 首先检查是否有RESTful风格的动作处理器
	if restAction, exists := t.RestActions[actionType]; exists {
		var result interface{}
		var err error

		// 根据HTTP方法调用对应的处理函数
		switch ctx.Request.Method {
		case http.MethodGet:
			result, err = restAction.Get(ctx, args)
		case http.MethodPost:
			result, err = restAction.Post(ctx, args)
		case http.MethodPut:
			result, err = restAction.Put(ctx, args)
		case http.MethodDelete:
			result, err = restAction.Delete(ctx, args)
		default:
			ctx.JSON(http.StatusMethodNotAllowed, gin.H{"error": "method not allowed for this action"})
			return
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, result)
		return
	}

	// 如果没有RESTful处理器，则使用普通处理函数
	handler, exists := t.Handlers[actionType]
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unsupported action type"})
		return
	}

	// 执行动作
	result, err := handler(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

package mgin

// Response 是一个使用泛型的响应结构体
type Response[T any] struct {
	Code int    `json:"code,omitempty" `
	Msg  string `json:"msg,omitempty"`
	Err  string `json:"err,omitempty"`
	Ver  string `json:"ver,omitempty"`
	Data T      `json:"data,omitempty"`
}

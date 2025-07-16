package zerr

import "fmt"

type Zerror struct {
	code   ErrorType
	msg    string
	detail string
}

// Error 实现官方error接口
func (z *Zerror) Error() string {
	return fmt.Sprintf("error code: %s, error message: %s, error detail: %s", z.code, z.msg, z.detail)
}

// NewZerror 新建错误
func (z *Zerror) NewZerror(code ErrorType, message, detail string) *Zerror {
	return &Zerror{
		code:   code,
		msg:    message,
		detail: detail,
	}
}

// Code 返回错误码
func (z *Zerror) Code() ErrorType {
	return z.code
}

// Msg 返回错误消息
func (z *Zerror) Msg() string {
	return z.msg
}

// Detail 返回错误详情信息
func (z *Zerror) Detail() string {
	return z.detail
}

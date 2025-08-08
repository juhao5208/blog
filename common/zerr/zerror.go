package zerr

import "fmt"

type Zerror struct {
	code ErrorType
	msg  string
}

// Error 实现官方error接口
func (z *Zerror) Error() string {
	return fmt.Sprintf("error code: %s, error message: %s, error detail: %s", z.code, z.msg)
}

// NewZerror 新建错误
func NewZerror(code ErrorType, message string) *Zerror {
	return &Zerror{
		code: code,
		msg:  message,
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

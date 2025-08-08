package zerr

type ErrorType string

const (
	InternalError ErrorType = "ZERROR2025071001" // 内部错误
	RunCmdError   ErrorType = "ZERROR2025071002" // 执行cmd错误
)

var ErrorMsgMap = map[ErrorType]string{
	InternalError: "服务内部错误",
	RunCmdError:   "运行命令%s错误:%s",
}

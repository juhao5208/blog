package zerr

type ErrorType string

const (
	InternalError ErrorType = "ZERROR2025071001" // 内部错误
	RunCmdError   ErrorType = "ZERROR2025071002" // 执行cmd错误
)

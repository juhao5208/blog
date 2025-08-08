package consts

const (
	LogFileName  = "blog.log" // 日志文件名
	LogFilePath  = ""         // 日志文件路径
	LogMaxSize   = 50         // 日志文件最大大小，单位MB
	LogMaxBackup = 5          // 日志文件最大备份数量
	LogLevel     = "info"     // 日志级别
	LogCompress  = false      // 日志是否压缩
	LogToStdout  = false      // 日志是否输出到标准输出
	LogIsJson    = false      // 日志是否为JSON格式
)

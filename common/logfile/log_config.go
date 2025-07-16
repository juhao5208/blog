package logfile

type LogConfig struct {
	Config *LogFile
}

type LogFile struct {
	FileName string `yaml:"file_name" validate:"required"`
	FilePath string `yaml:"file_path" validate:"required"`
	MaxSize  int    `yaml:"max_size" validate:"min=1 max=10"`
	LogLevel string `yaml:"log_level"`
	Compress bool   `yaml:"compress" validate:"oneof=trace debug info warn error fatal"`
	ToStdout bool   `yaml:"to_stdout"`
	IsJson   bool   `yaml:"is_json"`
}

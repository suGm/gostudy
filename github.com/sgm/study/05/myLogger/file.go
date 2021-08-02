package myLogger

type FileLogger struct {
	Level       LogLevel
	filePath    string // 文件路径
	fileName    string // 文件名
	maxFileSize int64  // 文件最大大小
}

// 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
}

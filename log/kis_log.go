package log

import "context"

// Logger抽象接口, 用于定义日志接口
type KisLogger interface {
	// InfoFX 有上下文的Info级别日志接口, format字符串格式
	InfoFX(ctx context.Context, str string, v ...interface{})
	// ErrorFX 有上下文的Error级别日志接口, format字符串格式
	ErrorFX(ctx context.Context, str string, v ...interface{})
	// DebugFX 有上下文的Debug级别日志接口, format字符串格式
	DebugFX(ctx context.Context, str string, v ...interface{})

	// InfoF 无上下文Info级别日志接口, format字符串格式
	InfoF(str string, v ...interface{})
	// ErrorF 无上下文Error级别日志接口, format字符串格式
	ErrorF(str string, v ...interface{})
	// DebugF 无上下文Debug级别日志接口, format字符串格式
	DebugF(str string, v ...interface{})
}

// kislog的默认kislog对象
var KisLog KisLogger

// SetLogger 设置kislog对象, 可以是用户自定义的kislog对象
func SetLogger(newlog KisLogger) {
	KisLog = newlog
}

// Logger 获取kislog对象
func Logger() KisLogger {
	return KisLog
}

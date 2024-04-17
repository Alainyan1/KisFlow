package log

import (
	"context"
	"fmt"
)

// KisDefaultLog默认日志对象
type KisDefaultLog struct{}

func (l *KisDefaultLog) InfoF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (l *KisDefaultLog) ErrorF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (l *KisDefaultLog) DebugF(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func (l *KisDefaultLog) InfoFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (l *KisDefaultLog) ErrorFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func (l *KisDefaultLog) DebugFX(ctx context.Context, str string, v ...interface{}) {
	fmt.Println(ctx)
	fmt.Printf(str, v...)
}

func init() {
	// 如果没有设置日志对象, 则设置默认日志对象
	if Logger() == nil {
		SetLogger(&KisDefaultLog{})
	}
}

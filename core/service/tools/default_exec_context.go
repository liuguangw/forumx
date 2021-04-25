package tools

import (
	"context"
	"time"
)

//DefaultExecContext 默认的执行超时设置(只在Controller内调用)
func DefaultExecContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.TODO(), 5*time.Second)
}

package tools

import (
	"math/rand"
	"time"
)

//init 初始化随机种子
func init() {
	rand.Seed(time.Now().UnixNano())
}

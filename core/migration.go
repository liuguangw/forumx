package core

//数据迁移接口定义
type Migration interface {
	//迁移的名称
	Name() string

	//执行迁移
	Up() error

	//回滚迁移
	Down() error
}

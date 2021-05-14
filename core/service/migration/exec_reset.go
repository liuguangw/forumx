package migration

import (
	"fmt"
)

//ExecuteReset 回滚所有已经执行了的迁移
func ExecuteReset() error {
	migrator := newMigrator()
	rollbackCount, err := migrator.ExecReset()
	if err != nil {
		return err
	}
	if rollbackCount == 0 {
		fmt.Println("nothing to rollback")
	} else {
		fmt.Println("rollback success: ", rollbackCount, " logs")
	}
	return nil
}

package migration

import (
	"fmt"
)

//ExecuteRollback 执行回滚
func ExecuteRollback(step int) error {
	migrator := newMigrator()
	rollbackCount, err := migrator.ExecRollback(step)
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

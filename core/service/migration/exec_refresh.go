package migration

import (
	"fmt"
)

//ExecuteRefresh 回滚所有已经执行了的迁移,再执行所有迁移
func ExecuteRefresh() error {
	migrator := newMigrator()
	rollbackCount, err := migrator.ExecReset()
	if err != nil {
		return err
	}
	if rollbackCount != 0 {
		fmt.Println("rollback success: ", rollbackCount, " logs")
	}
	migrateCount, err := migrator.ExecMigrate(0)
	if err != nil {
		return err
	}
	if migrateCount == 0 {
		fmt.Println("nothing to migrate")
	} else {
		fmt.Println("migrate success: ", migrateCount, " logs")
	}
	return nil
}

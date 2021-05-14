package migration

import (
	"fmt"
)

//ExecuteMigrate 执行迁移
func ExecuteMigrate(step int) error {
	migrator := newMigrator()
	migrateCount, err := migrator.ExecMigrate(step)
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

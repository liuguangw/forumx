package migration

import (
	"fmt"
	"github.com/liuguangw/migration"
)

//newMigrator 创建一个迁移工具
func newMigrator() *migration.Migrator {
	logHandler := &installedLogHandler{}
	return &migration.Migrator{
		InstalledLogHandler: logHandler,
		AllMigrations:       allMigrations,
		MigrateOutputHandler: func(name string) {
			fmt.Println("migrate " + name + " success")
		},
		RollbackOutputHandler: func(name string) {
			fmt.Println("rollback " + name + " success")
		},
	}
}

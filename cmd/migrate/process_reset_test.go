package migrate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessReset(t *testing.T) {
	migrationList := []Migration{
		&MockMigration{name: "a1"},
		&MockMigration{name: "a2"},
		&MockMigration{name: "a3"},
		&MockMigration{name: "a4"},
		&MockMigration{name: "a5"},
	}
	assertTool := assert.New(t)
	var (
		installedMigrationLogs []*migrationLog //已执行了的迁移记录
		rollbackNames          []string        //本次回滚的名称数组
	)
	migrationHandler := func(name string) error {
		rollbackNames = append(rollbackNames, name)
		return nil
	}
	installedMigrationLogs = []*migrationLog{
		{
			ID:    4,
			Name:  "a4",
			Batch: 3,
		},
		{
			ID:    3,
			Name:  "a3",
			Batch: 2,
		},
		{
			ID:    2,
			Name:  "a2",
			Batch: 2,
		},
		{
			ID:    1,
			Name:  "a1",
			Batch: 1,
		},
	}
	migrateCount, err := processRollback(installedMigrationLogs, migrationList, migrationHandler, 0)
	assertTool.Nil(err)
	tNames := []string{
		"a4", "a3", "a2", "a1",
	}
	assertTool.Equal(len(rollbackNames), migrateCount)
	for i, migratedName := range rollbackNames {
		assertTool.Equal(tNames[i], migratedName)
	}
}

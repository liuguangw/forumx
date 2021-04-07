package migrate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessRollback(t *testing.T) {
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
			Id:    5,
			Name:  "a5",
			Batch: 3,
		},
		{
			Id:    4,
			Name:  "a4",
			Batch: 3,
		},
		{
			Id:    3,
			Name:  "a3",
			Batch: 2,
		},
		{
			Id:    2,
			Name:  "a2",
			Batch: 2,
		},
		{
			Id:    1,
			Name:  "a1",
			Batch: 1,
		},
	}
	migrateCount, err := processRollback(installedMigrationLogs, migrationList, migrationHandler, 0)
	assertTool.Nil(err)
	tNames := []string{
		"a5", "a4",
	}
	assertTool.Equal(len(rollbackNames), migrateCount)
	for i, migratedName := range rollbackNames {
		assertTool.Equal(tNames[i], migratedName)
	}
	//step限制
	rollbackNames = nil
	migrateCount, err = processMigrate(installedMigrationLogs, migrationList, migrationHandler, 3)
	assertTool.Nil(err)
	tNames = []string{
		"a5", "a4", "a3",
	}
	assertTool.Equal(len(rollbackNames), migrateCount)
	for i, migratedName := range rollbackNames {
		assertTool.Equal(tNames[i], migratedName)
	}
}

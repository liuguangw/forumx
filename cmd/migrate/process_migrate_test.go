package migrate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockMigration struct {
	name string //迁移的名称
}

func (m *MockMigration) Name() string {
	return m.name
}

func (*MockMigration) Up() error {
	return nil
}

func (*MockMigration) Down() error {
	return nil
}

func TestProcessMigrate(t *testing.T) {
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
		migratedNames          []string        //本次迁移的名称数组
	)
	migrationHandler := func(name string) error {
		migratedNames = append(migratedNames, name)
		return nil
	}
	migrateCount, err := processMigrate(installedMigrationLogs, migrationList, migrationHandler, 0)
	assertTool.Nil(err)
	migrationNames := []string{
		"a1", "a2", "a3", "a4", "a5",
	}
	assertTool.Equal(len(migratedNames), migrateCount)
	for i, migratedName := range migratedNames {
		assertTool.Equal(migrationNames[i], migratedName)
	}
	//step限制
	migratedNames = nil
	migrateCount, err = processMigrate(installedMigrationLogs, migrationList, migrationHandler, 2)
	assertTool.Nil(err)
	migrationNames = []string{
		"a1", "a2",
	}
	assertTool.Equal(len(migratedNames), migrateCount)
	for i, migratedName := range migratedNames {
		assertTool.Equal(migrationNames[i], migratedName)
	}
	//已经执行了两条的情况
	installedMigrationLogs = []*migrationLog{
		{
			Name:  "a1",
			Batch: 1,
		},
		{
			Name:  "a2",
			Batch: 2,
		},
	}
	migratedNames = nil
	migrateCount, err = processMigrate(installedMigrationLogs, migrationList, migrationHandler, 2)
	assertTool.Nil(err)
	migrationNames = []string{
		"a3", "a4",
	}
	assertTool.Equal(len(migratedNames), migrateCount)
	for i, migratedName := range migratedNames {
		assertTool.Equal(migrationNames[i], migratedName)
	}
	//已经执行了a1,a3的情况
	installedMigrationLogs = []*migrationLog{
		{
			Name:  "a1",
			Batch: 1,
		},
		{
			Name:  "a3",
			Batch: 2,
		},
	}
	migratedNames = nil
	migrateCount, err = processMigrate(installedMigrationLogs, migrationList, migrationHandler, 2)
	assertTool.Nil(err)
	migrationNames = []string{
		"a2", "a4",
	}
	assertTool.Equal(len(migratedNames), migrateCount)
	for i, migratedName := range migratedNames {
		assertTool.Equal(migrationNames[i], migratedName)
	}
}

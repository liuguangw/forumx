package migrate

import (
	"github.com/liuguangw/forumx/core"
	"github.com/stretchr/testify/assert"
	"testing"
)

type A1Migration struct{}
type A2Migration struct{}
type A3Migration struct{}
type A4Migration struct{}
type A5Migration struct{}

func (a A1Migration) Name() string {
	return "a1"
}

func (a A1Migration) Up() error {
	println("a1_up")
	return nil
}

func (a A1Migration) Down() error {
	println("a1_down")
	return nil
}

func (a A2Migration) Name() string {
	return "a2"
}

func (a A2Migration) Up() error {
	println("a2_up")
	return nil
}

func (a A2Migration) Down() error {
	println("a2_down")
	return nil
}

func (a A3Migration) Name() string {
	return "a3"
}

func (a A3Migration) Up() error {
	println("a3_up")
	return nil
}

func (a A3Migration) Down() error {
	println("a3_down")
	return nil
}

func (a A4Migration) Name() string {
	return "a4"
}

func (a A4Migration) Up() error {
	println("a4_up")
	return nil
}

func (a A4Migration) Down() error {
	println("a4_down")
	return nil
}

func (a A5Migration) Name() string {
	return "a5"
}

func (a A5Migration) Up() error {
	println("a5_up")
	return nil
}

func (a A5Migration) Down() error {
	println("a5_down")
	return nil
}

var migrations = []core.Migration{
	&A1Migration{},
	&A2Migration{},
	&A3Migration{},
	&A4Migration{},
	&A5Migration{},
}

func TestProcessMigrate(t *testing.T) {
	assertTool := assert.New(t)
	var installedMigrationLogs []*installedMigrationLog
	println("===============")
	migrationLogs, err := processMigrate(installedMigrationLogs, migrations, 0)
	assertTool.Nil(err)
	migrationNames := []string{
		"a1", "a2", "a3", "a4", "a5",
	}
	for i, migrationLog := range migrationLogs {
		assertTool.Equal(migrationNames[i], migrationLog.Name)
	}
	//step限制
	println("===============")
	migrationLogs, err = processMigrate(installedMigrationLogs, migrations, 2)
	assertTool.Nil(err)
	migrationNames = []string{
		"a1", "a2",
	}
	for i, migrationLog := range migrationLogs {
		assertTool.Equal(migrationNames[i], migrationLog.Name)
	}
	//已经执行了两条的情况
	installedMigrationLogs = []*installedMigrationLog{
		{
			Name:  "a1",
			Batch: 1,
		},
		{
			Name:  "a2",
			Batch: 2,
		},
	}
	println("===============")
	migrationLogs, err = processMigrate(installedMigrationLogs, migrations, 2)
	assertTool.Nil(err)
	migrationNames = []string{
		"a3", "a4",
	}
	for i, migrationLog := range migrationLogs {
		assertTool.Equal(migrationNames[i], migrationLog.Name)
		assertTool.Equal(3, migrationLog.Batch)
	}
	//已经执行了a1,a3的情况
	installedMigrationLogs = []*installedMigrationLog{
		{
			Name:  "a1",
			Batch: 1,
		},
		{
			Name:  "a3",
			Batch: 2,
		},
	}
	println("===============")
	migrationLogs, err = processMigrate(installedMigrationLogs, migrations, 2)
	assertTool.Nil(err)
	migrationNames = []string{
		"a2", "a4",
	}
	for i, migrationLog := range migrationLogs {
		assertTool.Equal(migrationNames[i], migrationLog.Name)
		assertTool.Equal(3, migrationLog.Batch)
	}
}

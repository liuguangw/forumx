package migration

import (
	"github.com/liuguangw/forumx/core/migrations"
	"github.com/liuguangw/migration"
)

//获取应用应该执行的所有迁移对象
func allMigrations() []migration.Migration {
	migrationList := []migration.Migration{
		new(migrations.CreateCountersCollection),
		new(migrations.CreateUsersCollection),
		new(migrations.CreateUserEmailLinksCollection),
		new(migrations.CreateUsersEmailsCollection),
		new(migrations.CreateUserMobileCodesCollection),
		new(migrations.CreateUserMobilesCollection),
		new(migrations.InitCountersCollection),
		new(migrations.CreateUserSessionsCollection),
		new(migrations.CreateUserTotpKeysCollection),
		new(migrations.CreateCachesCollection),
		new(migrations.CreateAppConfigsCollection),
		new(migrations.InitAppConfigsCollection),
		new(migrations.CreateForumAreasCollection),
		new(migrations.CreateForumsCollection),
		new(migrations.InitForumAreasCollection),
		new(migrations.InitForumsCollection),
	}
	return migrationList
}

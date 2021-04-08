package migrate

import "github.com/liuguangw/forumx/core/migrations"

//获取应用应该执行的所有迁移对象
func allMigrations() []Migration {
	migrationList := []Migration{
		new(migrations.CreateCountersCollection),
		new(migrations.CreateUsersCollection),
		new(migrations.CreateUserEmailLinksCollection),
		new(migrations.CreateUsersEmailsCollection),
		new(migrations.CreateUserMobileCodesCollection),
		new(migrations.CreateUserMobilesCollection),
		new(migrations.InitCountersCollection),
	}
	return migrationList
}

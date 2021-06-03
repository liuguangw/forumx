package migration

//已执行的数据迁移记录
type migrationLog struct {
	ID    int    `bson:"id"`    //迁移的ID
	Name  string `bson:"name"`  //迁移名称
	Batch int    `bson:"batch"` //迁移的批次
}

package test_helpers

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func ClearUserTable() {
	o := orm.NewOrm()
	_, err := o.Raw("TRUNCATE TABLE \"user\"").Exec()
	if err != nil {
		// If table can't be truncated, rebuild all tables (CAUTION: Star and Message db are lost!)
		// This is only for absolute startup
		ResetDB()
	}
}

func ResetDB() {
	// If table can't be truncated, rebuild all tables (CAUTION: Star and Message db are lost!)
	// This is only for absolute startup
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		logs.Error(err)
	}
}

package test_helpers

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func TruncateTables(tables ...string) {
	o := orm.NewOrm()
	rawSQL := ""

	for _, table := range tables {
		rawSQL += fmt.Sprintf("TRUNCATE TABLE \"%s\";", table)
	}

	if rawSQL == "" {
		resetDB()
		return
	}

	_, err := o.Raw(rawSQL).Exec()
	if err != nil {
		// If tables can't be truncated, rebuild all tables
		resetDB()
	}
}

func resetDB() {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		logs.Error(err)
	}
}

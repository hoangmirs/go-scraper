package bootstrap

import (
	"fmt"

	_ "github.com/hoangmirs/go-scraper/models" // Models

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq" // Postgres driver
)

func SetUpDB() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		logs.Critical(fmt.Sprintf("Failed to register the driver %v", err))
	}

	err = orm.RegisterDataBase("default", "postgres", web.AppConfig.DefaultString("dbUrl", ""))
	if err != nil {
		logs.Critical(fmt.Sprintf("Failed to connect to the database %v", err))
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		logs.Critical(fmt.Sprintf("Failed to sync the database %v", err))
	}

	if web.AppConfig.DefaultString("runmode", "dev") == "dev" {
		orm.Debug = true
	}
}

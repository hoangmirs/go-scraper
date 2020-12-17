package bootstrap

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq" // Postgres driver
)

func SetUpDB() {
	err := orm.RegisterDriver("postgres", orm.DRPostgres)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to register the driver %v", err))
	}

	err = orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("dbUrl"))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to the database %v", err))
	}

	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to sync the database %v", err))
	}
}

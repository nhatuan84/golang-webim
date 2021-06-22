package main

import (
	_ "gochat/routers"
	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/client/orm"
 	_ "github.com/lib/pq"
	"gochat/models"
)

func init() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", 
    					"postgres", 
    					"host=127.0.0.1 port=5432 user=postgres password=xxx dbname=test_orm sslmode=disable")
    orm.RegisterModel(new(models.User))
    orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Run()
}


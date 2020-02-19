package main

import (
    _ "github.com/GoAdminGroup/go-admin/adapter/gin"
    _ "github.com/GoAdminGroup/themes/adminlte"
    _ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"

    "github.com/GoAdminGroup/go-admin/engine"
    // "github.com/GoAdminGroup/go-admin/examples/datamodel"
    "github.com/GoAdminGroup/go-admin/modules/config"
    "github.com/GoAdminGroup/go-admin/modules/language"
    "github.com/GoAdminGroup/go-admin/plugins/admin"
    // "github.com/GoAdminGroup/go-admin/template/types"
    "github.com/gin-gonic/gin"

    "meeting-admin/tables"
)

func main() {
    cfg := config.Config{
        Databases: config.DatabaseList{
            "default": {
                Host:       "127.0.0.1",
                Port:       "3306",
                User:       "root",
                Pwd:        "root1234",
                Name:       "meeting",
                MaxIdleCon: 50,
                MaxOpenCon: 150,
                Driver:     config.DriverMysql,
            },
        },
        UrlPrefix: "admin",
        Store: config.Store{
            Path:   "./uploads",
            Prefix: "uploads",
        },
        Language: language.CN,
    }
    r := gin.Default()
    eng := engine.Default()
    adminPlugin := admin.NewAdmin(tables.Generators)
    if err := eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r); err != nil {
        panic(err)
    }
    _ = r.Run(":9033")
}
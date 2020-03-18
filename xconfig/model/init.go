package model

import (
	"database/sql"
	"base/pkg/xmysql"
)

var (
	ConfigModel	*configModel

	dbHandle	*sql.DB
)

func init()  {
	var err error
	ConfigModel = new(configModel)

	//init handle
	if dbHandle,err = xmysql.NewMysql("db.local");err != nil {
		panic(err)
	}
}
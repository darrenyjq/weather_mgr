package model

import (
	"database/sql"
	"fmt"
)

type configModel struct {

}

func (this configModel) GetConfig(key string, group int64, appName string) (jsonVal []byte,err error) {
	row := dbHandle.QueryRow("select `config` from configs where `key`=? and `group`=? and `app_name`=? and `status`=0 limit 1",key,group,appName)
	var val string
	err = row.Scan(&val)
	if err == sql.ErrNoRows {
		err = fmt.Errorf("key:%s group:%d os:%s not exist in db",key,group,appName)
		return
	}
	if err != nil {
		return
	}

	jsonVal = []byte(val)
	return
}


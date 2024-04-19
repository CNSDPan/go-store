package test

import (
	"store/yaml"
	"testing"
)

var db = gorm_db.GetReadDB(yaml.ReadMysqlCon.Name)

func TestInitMysql(t *testing.T) {

	//fmt.Printf("db :%v", db)
}

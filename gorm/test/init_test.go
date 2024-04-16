package test

import (
	"store/gorm"
	"store/yaml"
	"testing"
)

var db = gorm.GetReadDB(yaml.ReadMysqlCon.Name)

func TestInitMysql(t *testing.T) {

	//fmt.Printf("db :%v", db)
}

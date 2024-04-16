package yaml

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
)

type MysqlConf struct {
	Name       string       `json:",optional"`
	Separation int          `json:",optional"`
	MasterDB   MasterDBConf `json:",optional,inherit"`
	SlaveDB    SlaveDBConf  `json:",optional,inherit"`
	Charset    string       `json:",optional"`
}
type MasterDBConf struct {
	Host     string `json:",optional"`
	Port     string `json:",optional"`
	User     string `json:",optional"`
	Password string `json:",optional"`
	Database string `json:",optional"`
}
type SlaveDBConf struct {
	Host     []string `json:",optional"`
	Port     string   `json:",optional"`
	User     string   `json:",optional"`
	Password string   `json:",optional"`
	Database string   `json:",optional"`
}

var ReadMysqlCon *MysqlConf

const (
	SEPARATION_YES = 1
	SEPARATION_NO  = 2
)

var separationName = map[int]string{
	SEPARATION_YES: "读写分离",
	SEPARATION_NO:  "单个数据库实例",
}

func init() {
	// 获取配置文件的路径
	realPath := getCurrentDir()
	mysqlFilePath := realPath + "/mysql.yaml"
	mysqlFile := flag.String("mysql-f", mysqlFilePath, "the mysql config file")

	var c MysqlConf
	conf.MustLoad(*mysqlFile, &c)
	ReadMysqlCon = &c
}

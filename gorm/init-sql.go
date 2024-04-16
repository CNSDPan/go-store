package gorm

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"store/yaml"
	"time"
)

var dbReadMysql = map[string]*gorm.DB{}

func init() {
	c := yaml.ReadMysqlCon
	ctx := context.Background()
	logg := logx.WithContext(ctx)
	logg.Infof("%s mysql connect db init...", c.Name)
	switch c.Separation {
	case yaml.SEPARATION_YES:
		initRWDB(c)
	case yaml.SEPARATION_NO:
		initDB(c)
	default:
		panic("Separation undulated ")
	}
	logg.Infof("%s mysql connect db init ok", c.Name)
}

// initDB
// @Auth：parker
// @Desc：单个数据库实例
// @Date：2024-04-15 17:15:45
// @param：c
func initDB(c *yaml.MysqlConf) {
	var dbConn *gorm.DB
	var err error
	dsn := c.MasterDB.User + ":" + c.MasterDB.Password + "@tcp(" + c.MasterDB.Host + ":" + c.MasterDB.Port + ")/" + c.MasterDB.Database +
		"?loc=Local&parseTime=True&charset=" + c.Charset
	dbConn, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("单个数据库实例初始化失败")
	}
	db, e := dbConn.DB()
	if e != nil {
		panic("单个数据库实例: 获取 数据库 实例失败")
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	db.SetConnMaxLifetime(time.Hour * 2)
	dbReadMysql[c.Name] = dbConn
}

func initRWDB(c *yaml.MysqlConf) {

}

// GetReadDB
// @Auth：parker
// @Desc：获取连接池的DB
// @Date：2024-04-15 13:41:29
// @param：dbName
// @return：db | nil
func GetReadDB(dbName string) (db *gorm.DB) {
	var ok bool
	if db, ok = dbReadMysql[dbName]; ok {
		return db
	} else {
		return nil
	}
}

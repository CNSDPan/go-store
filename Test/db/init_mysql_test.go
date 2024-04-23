package db

import (
	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
	"store/gorm_db"
	"store/yaml"
	"strconv"
)

var ServiceIdRpc = "199"
var DBModel *struct {
	DB   *gorm.DB
	Node *snowflake.Node
}

func init() {
	var err error
	var nodeId int64
	var node *snowflake.Node
	var db = gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	nodeId, err = strconv.ParseInt(ServiceIdRpc, 0, 64)
	if err != nil {
		panic("转换 64位整形失败 " + err.Error())
	}
	node, err = snowflake.NewNode(nodeId)
	if err != nil {
		panic("new 节点失败 " + err.Error())
	}
	DBModel = &struct {
		DB   *gorm.DB
		Node *snowflake.Node
	}{DB: db, Node: node}
}

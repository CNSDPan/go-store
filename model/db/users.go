package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
)

const (
	USER_STATUS_1 int8 = 1
	USER_STATUS_2 int8 = 2
)

var StatusName = map[int8]string{
	USER_STATUS_1: "启用",
	USER_STATUS_2: "禁用",
}

type Users struct {
	ID        uint32 `gorm:"primaryKey;column:id" json:"-"`
	UserID    int64  `gorm:"column:user_id" json:"userId"`       // 用户IID
	Token     string `gorm:"column:token" json:"token"`          // token
	Status    int8   `gorm:"column:status" json:"status"`        // 1=启用 2=禁用
	Name      string `gorm:"column:name" json:"name"`            // 昵称
	Fund      int64  `gorm:"column:fund" json:"fund"`            // 用户资金,入库*1000【1000 = 1元】
	CreatedAt string `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt string `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type _UsersMgr struct {
	*_BaseMgr
}

func UsersTableName() string {
	return "users"
}

// UsersMgr open func
func UsersMgr() *_UsersMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table(UsersTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Reset 重置gorm会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UsersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Users, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

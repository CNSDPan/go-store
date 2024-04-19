package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
)

type ChinaArea struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`
	Code     uint64 `gorm:"column:code" json:"code"`         // 区划代码
	Name     string `gorm:"column:name" json:"name"`         // 名称
	Level    bool   `gorm:"column:level" json:"level"`       // 级别1-5,省市县镇村
	Pcode    int64  `gorm:"column:pcode" json:"pcode"`       // 父级区划代码
	Category int    `gorm:"column:category" json:"category"` // 城乡分类
}

type _ChinaAreaMgr struct {
	*_BaseMgr
}

func ChinaAreaTableName() string {
	return "china_area"
}

func ChinaAreaMgr() *_ChinaAreaMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_ChinaAreaMgr{_BaseMgr: &_BaseMgr{DB: db.Table(ChinaAreaTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ChinaAreaMgr) GetTableName() string {
	return "china_area"
}

// Reset 重置gorm会话
func (obj *_ChinaAreaMgr) Reset() *_ChinaAreaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ChinaAreaMgr) Get() (result ChinaArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ChinaArea{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ChinaAreaMgr) Gets() (results []*ChinaArea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(ChinaArea{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ChinaAreaMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]ChinaArea, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(ChinaArea{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
	"time"
)

type FundIncome struct {
	ID        uint32    `gorm:"primaryKey;column:id" json:"-"`
	IncomeID  int64     `gorm:"column:income_id" json:"incomeId"`   // 收入IID
	UserID    int64     `gorm:"column:user_id" json:"userId"`       // 用户IID
	Type      bool      `gorm:"column:type" json:"type"`            // 收入类型：2-红包
	Before    uint64    `gorm:"column:before" json:"before"`        // 收入前
	After     uint64    `gorm:"column:after" json:"after"`          // 收入后
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type _FundIncomeMgr struct {
	*_BaseMgr
}

func FundIncomeTableName() string {
	return "fund_income"
}

// FundIncomeMgr open func
func FundIncomeMgr() *_FundIncomeMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_FundIncomeMgr{_BaseMgr: &_BaseMgr{DB: db.Table(FundIncomeTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FundIncomeMgr) GetTableName() string {
	return "fund_income"
}

// Reset 重置gorm会话
func (obj *_FundIncomeMgr) Reset() *_FundIncomeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FundIncomeMgr) Get() (result FundIncome, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FundIncome{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FundIncomeMgr) Gets() (results []*FundIncome, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FundIncome{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FundIncomeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FundIncome, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FundIncome{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

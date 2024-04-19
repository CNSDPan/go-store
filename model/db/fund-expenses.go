package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
	"time"
)

type FundExpenses struct {
	ID         uint32    `gorm:"primaryKey;column:id" json:"-"`
	ExpensesID int64     `gorm:"column:expenses_id" json:"expensesId"` // 支出IID
	UserID     int64     `gorm:"column:user_id" json:"userId"`         // 用户IID
	Type       bool      `gorm:"column:type" json:"type"`              // 支出类型：1-订单、2-红包
	Before     uint64    `gorm:"column:before" json:"before"`          // 扣除前
	After      uint64    `gorm:"column:after" json:"after"`            // 扣除后
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`   // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`   // 更新时间
}

type _FundExpensesMgr struct {
	*_BaseMgr
}

func FundExpensesTableName() string {
	return "fund_expenses"
}

// FundExpensesMgr open func
func FundExpensesMgr() *_FundExpensesMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_FundExpensesMgr{_BaseMgr: &_BaseMgr{DB: db.Table(FundExpensesTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_FundExpensesMgr) GetTableName() string {
	return "fund_expenses"
}

// Reset 重置gorm会话
func (obj *_FundExpensesMgr) Reset() *_FundExpensesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_FundExpensesMgr) Get() (result FundExpenses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FundExpenses{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_FundExpensesMgr) Gets() (results []*FundExpenses, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(FundExpenses{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_FundExpensesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]FundExpenses, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(FundExpenses{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

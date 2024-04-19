package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
	"time"
)

type PayOrder struct {
	ID        uint32    `gorm:"primaryKey;column:id" json:"-"`
	PayID     int64     `gorm:"column:pay_id" json:"payId"`         // 支付单编号IID
	OrderID   int64     `gorm:"column:order_id" json:"orderId"`     // 订单编号IID
	Type      bool      `gorm:"column:type" json:"type"`            // 支付来源：1-微信、2-支付宝、3-美团、4-第三方
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type _PayOrderMgr struct {
	*_BaseMgr
}

func PayOrderTableName() string {
	return "pay_order"
}

// PayOrderMgr open func
func PayOrderMgr() *_PayOrderMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_PayOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table(PayOrderTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_PayOrderMgr) GetTableName() string {
	return "pay_order"
}

// Reset 重置gorm会话
func (obj *_PayOrderMgr) Reset() *_PayOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_PayOrderMgr) Get() (result PayOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PayOrder{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_PayOrderMgr) Gets() (results []*PayOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(PayOrder{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_PayOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]PayOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(PayOrder{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

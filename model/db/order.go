package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
	"time"
)

type Order struct {
	ID              uint32    `gorm:"primaryKey;column:id" json:"-"`
	OrderID         int64     `gorm:"column:order_id" json:"orderId"`                 // 订单编号
	Status          uint8     `gorm:"column:status" json:"status"`                    // 订单状态：1-待支付、11-取消订单、12-失效订单、20-已完成
	PayStatus       uint8     `gorm:"column:pay_status" json:"payStatus"`             // 支付状态:1-待支付、2-取消支付、3-支付超时、11-支付失败、20-已支付
	PayTime         int64     `gorm:"column:pay_time" json:"payTime"`                 // 支付时间,毫秒
	PayTimeout      int64     `gorm:"column:pay_timeout" json:"payTimeout"`           // 支付有效时间,毫秒
	PayTimeClose    int64     `gorm:"column:pay_time_close" json:"payTimeClose"`      // 支付“取消|失效 ”时间,毫秒
	Total           uint64    `gorm:"column:total" json:"total"`                      // 订单总价,入库*1000【1000 = 1元】
	Quantity        uint32    `gorm:"column:quantity" json:"quantity"`                // 商品总数量
	Remark          string    `gorm:"column:remark" json:"remark"`                    // 订单备注
	AddressName     string    `gorm:"column:address_name" json:"addressName"`         // 收货人姓名
	AddressPhone    string    `gorm:"column:address_phone" json:"addressPhone"`       // 收货人电话
	AddressCountry  int       `gorm:"column:address_country" json:"addressCountry"`   // 国家
	AddressProvince int       `gorm:"column:address_province" json:"addressProvince"` // 省
	AddressCity     int       `gorm:"column:address_city" json:"addressCity"`         // 市
	AddressDistrict int       `gorm:"column:address_district" json:"addressDistrict"` // 区
	AddressDetail   string    `gorm:"column:address_detail" json:"addressDetail"`     // 详细地址
	CreatedAt       time.Time `gorm:"column:created_at" json:"createdAt"`             // 创建时间
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updatedAt"`             // 更新时间
}

type _OrderMgr struct {
	*_BaseMgr
}

func OrderTableName() string {
	return "order"
}

// OrderMgr open func
func OrderMgr() *_OrderMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table(OrderTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderMgr) GetTableName() string {
	return "order"
}

// Reset 重置gorm会话
func (obj *_OrderMgr) Reset() *_OrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_OrderMgr) Get() (result Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Order{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderMgr) Gets() (results []*Order, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Order{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_OrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Order, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Order{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
	"time"
)

type OrderDetail struct {
	ID             uint32    `gorm:"primaryKey;column:id" json:"-"`
	DetailID       int64     `gorm:"column:detail_id" json:"detailId"`             // 明细编号IID
	OrderID        int64     `gorm:"column:order_id" json:"orderId"`               // 订单编号IID
	ProductID      int16     `gorm:"column:product_id" json:"productId"`           // 商品IID
	Total          uint64    `gorm:"column:total" json:"total"`                    // 商品总价入库*1000【1000 = 1元】
	Quantity       uint32    `gorm:"column:quantity" json:"quantity"`              // 商品数量
	Price          uint64    `gorm:"column:price" json:"price"`                    // 商品单价入库*1000【1000 = 1元】
	ReturnQuantity uint32    `gorm:"column:return_quantity" json:"returnQuantity"` // 退货数量
	Remark         string    `gorm:"column:remark" json:"remark"`                  // 备注
	CreatedAt      time.Time `gorm:"column:created_at" json:"createdAt"`           // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updatedAt"`           // 更新时间
}

type _OrderDetailMgr struct {
	*_BaseMgr
}

func OrderDetailTableName() string {
	return "order_detail"
}

// OrderDetailMgr open func
func OrderDetailMgr() *_OrderDetailMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_OrderDetailMgr{_BaseMgr: &_BaseMgr{DB: db.Table(OrderDetailTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_OrderDetailMgr) GetTableName() string {
	return "order_detail"
}

// Reset 重置gorm会话
func (obj *_OrderDetailMgr) Reset() *_OrderDetailMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_OrderDetailMgr) Get() (result OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderDetail{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_OrderDetailMgr) Gets() (results []*OrderDetail, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(OrderDetail{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_OrderDetailMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]OrderDetail, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(OrderDetail{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

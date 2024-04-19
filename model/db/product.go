package db

import (
	"context"
	"store/gorm_db"
	"store/yaml"
	"time"
)

type Product struct {
	ID        uint32    `gorm:"primaryKey;column:id" json:"-"`
	ProductID int64     `gorm:"column:product_id" json:"productId"` // 商品IID
	Title     string    `gorm:"column:title" json:"title"`          // 商品中文名称
	Image     string    `gorm:"column:image" json:"image"`          // 商品图片
	Status    bool      `gorm:"column:status" json:"status"`        // 状态:1-正在销售、2-新品、2-爆款、9-停止销售
	Price     uint64    `gorm:"column:price" json:"price"`          // 价钱
	CreatedAt time.Time `gorm:"column:created_at" json:"createdAt"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updatedAt"` // 更新时间
}

type _ProductMgr struct {
	*_BaseMgr
}

func ProductTableName() string {
	return "product"
}

// ProductMgr open func
func ProductMgr() *_ProductMgr {
	db := gorm_db.GetReadDB(yaml.RWMysqlCon.Name)
	ctx, cancel := context.WithCancel(context.Background())
	return &_ProductMgr{_BaseMgr: &_BaseMgr{DB: db.Table(ProductTableName()), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_ProductMgr) GetTableName() string {
	return "product"
}

// Reset 重置gorm会话
func (obj *_ProductMgr) Reset() *_ProductMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_ProductMgr) Get() (result Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Find(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_ProductMgr) Gets() (results []*Product, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Product{}).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_ProductMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Product, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Product{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

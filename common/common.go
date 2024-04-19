package common

import "github.com/shopspring/decimal"

// EnterExchange
// @Auth：parker
// @Desc：统一入库金额换算,将金额 * 10000
// @Date：2024-04-19 18:28:20
// @param：price 只支持 int | int8 | int32 | int64 | float32 | float64
// @return：dbPrice
func EnterExchange(price interface{}) (dbPrice int64) {
	var rate int64 = 1000
	switch price.(type) {
	case int:
		newPrice, _ := price.(int64)
		dbPrice = decimal.NewFromInt(newPrice).Mul(decimal.NewFromInt(rate)).IntPart()
	case int8:
		newPrice, _ := price.(int8)
		dbPrice = decimal.NewFromInt(int64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case int32:
		newPrice, _ := price.(int32)
		dbPrice = decimal.NewFromInt(int64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case int64:
		newPrice, _ := price.(int64)
		dbPrice = decimal.NewFromInt(newPrice).Mul(decimal.NewFromInt(rate)).IntPart()
	case float32:
		newPrice, _ := price.(float32)
		dbPrice = decimal.NewFromFloat(float64(newPrice)).Mul(decimal.NewFromInt(rate)).IntPart()
	case float64:
		newPrice, _ := price.(float64)
		dbPrice = decimal.NewFromFloat(newPrice).Mul(decimal.NewFromInt(rate)).IntPart()
	}
	return dbPrice
}

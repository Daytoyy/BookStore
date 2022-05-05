package model

type OrderItem struct{
	Count int64 //数量
	Amount float64 //金额小计
	Title string //书名
	Author string //作者
	Price float64 //单价
	ImgPath string //图书封面
	OrderId string //订单id
}
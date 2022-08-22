package model

type OrderItem struct {
	OrderItemID int64   //订单项的id
	Count       int64   //订单项的数量
	Amount      float64 //订单项中金额小计
	Title       string  //订单中的书名
	Author      string  //订单项中图书的作者
	Price       float64 //订单项中图书的单价
	ImgPath     string  //订单项中图书的图片
	OrderID     string  //订单项所属的订单
}

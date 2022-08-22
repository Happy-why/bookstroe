package model

type Order struct {
	OrderID     string  //订单号 唯一
	CreateTime  string  //生成订单的时间
	TotalCount  int64   //订单图书的总数量
	TotalAmount float64 //订单图书的总价格
	State       int64   //订单状态 0 未发货  1 已发货 2 交易完成
	UserID      int64   //用户ID
}

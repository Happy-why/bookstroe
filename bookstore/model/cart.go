package model

//Cart 购物车结构体
type Cart struct {
	CartID      string      //购物车的id
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物车中图书的总数量
	TotalAmount float64     //购物车中图书的总价值
	UserID      int         //当前购物车所属的车主
}

//GetTotalCount 获取图书的总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中购物项的切片
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//GetTotalAmount 获取图书的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var TotalAmount float64
	for _, v := range cart.CartItems {
		TotalAmount = TotalAmount + v.GetAmount()
	}
	return TotalAmount
}

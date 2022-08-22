package dao

import (
	"fmt"
	"github.com/model"
	"github.com/utils"
)

//AddCart 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//写sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//获取购物车中所有的购物项
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		err := AddCartItem(cartItem)
		fmt.Println("AddCart AddCartItem err:", err)
		if err != nil {
			return err
		}
	}
	return nil
}

//GetCartByUserID 根据用户的id从数据库中查询对应的购物车
func GetCartByUserID(userID int) *model.Cart {
	sqlStr := "select id,total_count,total_amount,user_id from carts where user_id=?"

	row := utils.Db.QueryRow(sqlStr, userID)
	//创建一个购物车
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil
	}
	//获取当前购物车中所有的购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	//将购物项添加到购物车中去
	cart.CartItems = cartItems
	return cart
}

//UpdateCart 更新购物车中总数量和总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count= ? , total_amount= ? where id = ? "
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartByCartID 根据购物车的id删除购物车
func DeleteCartByCartID(cartID string) error {
	//删除购物车之前需要先删除所有的购物项
	err := DeleteCartItemsByCartID(cartID)
	if err != nil {
		fmt.Println("DeleteCartByCartID DeleteCartItemsByCartID err:", err)
		return err
	}
	//写sql语句
	sqlStr := "delete from carts where id=?"
	_, err2 := utils.Db.Exec(sqlStr, cartID)
	if err2 != nil {
		fmt.Println("DeleteCartByCartID Db.Exec err:", err2)
		return err2
	}
	return nil
}

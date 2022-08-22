package dao

import (
	"github.com/model"
	"github.com/utils"
)

//AddCartItem 向购物项表中添加购物项
func AddCartItem(cartItem *model.CartItem) error {
	//写sql语句
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemByBookIDAndCartID 根据图书的id和购物车的id去获取对应的购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id=? and cart_id=?"
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	//根据图书的id查询图书信息
	book := GetBookById(bookID)
	//将book设置到购物项
	cartItem.Book = book
	return cartItem, nil
}

//UpdateBookCount 更新购物项的图书的数量
func UpdateBookCount(item *model.CartItem) error {
	sqlStr := "update cart_items set count=?,amount=? where cart_id=? and book_id=?"
	_, err := utils.Db.Exec(sqlStr, item.Count, item.GetAmount(), item.CartID, item.Book.ID)
	if err != nil {
		return err
	}
	return nil
}

//GetCartItemsByCartID 根据购物扯的id获取购物车中所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id=?"
	rows, _ := utils.Db.Query(sqlStr, cartID)
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量接收bookID
		var bookID string
		cartItem := &model.CartItem{}
		err := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err != nil {
			return nil, err
		}
		book := GetBookById(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//DeleteCartItemsByCartID 根据购物车的id删除所有的购物项
func DeleteCartItemsByCartID(cartID string) error {
	//写sql语句
	sqlStr := "delete from cart_items where cart_id=?"
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteCartItemByID 根据购物项的id删除购物项
func DeleteCartItemByID(cartItemID string) error {
	//写sql语句
	sqlStr := "delete from cart_items where id=?"
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	if err != nil {
		return err
	}
	return nil
}

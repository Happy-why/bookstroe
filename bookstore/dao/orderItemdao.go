package dao

import (
	"fmt"
	"github.com/model"
	"github.com/utils"
)

//AddOrderItem 添加订单项
func AddOrderItem(item *model.OrderItem) error {
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, item.Count, item.Amount, item.Title, item.Author, item.Price, item.ImgPath, item.OrderID)
	if err != nil {
		fmt.Println("AddOrderItem Db.Exec err:", err)
		return err
	}
	return nil
}

//GetOrderItemsByID 根据订单号获取所有的订单项
func GetOrderItemsByID(orderID string) ([]*model.OrderItem, error) {
	sqlStr := "select id,count,amount,title,author,price,img_path from order_items where order_id=?"
	rows, err := utils.Db.Query(sqlStr, orderID)
	if err != nil {
		fmt.Println("GetOrderItemsByID Db.Query err:", err)
		return nil, err
	}
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		err := rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath)
		if err != nil {
			fmt.Println("GetOrderItemsByID rows.Next Scan err:", err)
			return nil, err
		}
		//添加到切片中
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

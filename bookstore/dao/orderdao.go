package dao

import (
	"fmt"
	"github.com/model"
	"github.com/utils"
)

//AddOrder 向数据库中插入订单
func AddOrder(order *model.Order) error {
	//写sql语句
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}

//GetOrders 获取数据库中所有的订单
func GetOrders() ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state,user_id from orders"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println("GetOrders Db.Query err:", err)
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserID)
		if err != nil {
			fmt.Println("GetOrders rows.Next Scan err:", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

// GetMyOrder 获取我的订单
func GetMyOrder(userID string) ([]*model.Order, error) {
	sqlStr := "select id,create_time,total_count,total_amount,state from orders where user_id=?"
	rows, err := utils.Db.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		err := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State)
		if err != nil {
			fmt.Println(" GetMyOrder rows.Next Scan err:", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

//UpdateOrderState 更新订单的状态，即发货和收货
func UpdateOrderState(orderID string, state int64) error {
	sql := "update orders set state = ? where id = ?"
	_, err := utils.Db.Exec(sql, state, orderID)
	if err != nil {
		fmt.Println("UpdateOrderState Db.Exec err:", err)
		return err
	}
	return nil
}

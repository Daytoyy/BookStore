package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"strconv"
	"time"
)

//生成订单
func CreatOrder(userId int) *model.Order {

	//根据同户id获取该用户的购物车
	userid := strconv.Itoa(userId)
	cart, _, _ := GetCartByUserId(userid)
	timStr := time.Now().Format("2006-01-02 15:04:05")
	order := &model.Order{
		OrderId:     utils.CreateUUID(),
		CreateTime:  timStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserId:      int64(userId),
		CartId:      cart.CartId,
	}
	return order
}

func AddOder(order *model.Order) error {
	//将订单保存到数据库
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		fmt.Println("AddOder预编译和执行出错:", err)
		return err
	}
	return nil
}

//获取当前用户的所有订单
func GetOrders(userId int) ([]*model.Order, error) {
	//写sql语句
	sqlStr := "select id,create_time,total_count,total_amount,state from orders where user_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, userId)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		//给book中的字段赋值
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State)
		//将book添加到books中
		orders = append(orders, order)
	}
	return orders, err
}

//获取所有订单
func GetAllOrders() ([]*model.Order, error) {
	//写sql语句
	sqlStr := "select id,create_time,total_count,total_amount,state from orders"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		//给book中的字段赋值
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State)
		if err != nil {
			return nil, err
		}
		//将book添加到books中
		orders = append(orders, order)
	}
	return orders, err
}

//处理订单发货 根据订单id处理发货 更新订单状态
func SendOrderByOrderId(orderId string, state string) error {
	//写sql语句
	sqlStr := "update orders set state=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, state, orderId)
	if err != nil {
		return err
	}
	return nil
}

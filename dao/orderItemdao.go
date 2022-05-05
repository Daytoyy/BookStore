package dao

import (
	//"time"
	//"strconv"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//生成订单详情
func CreatOrderItem(cartItem *model.CartItem) *model.OrderItem {
	orderItem := &model.OrderItem{
		Count:   cartItem.Count,
		Amount:  cartItem.Amount,
		Title:   cartItem.Book.Title,
		Author:  cartItem.Book.Author,
		Price:   cartItem.Book.Price,
		ImgPath: cartItem.Book.ImgPath,
	}
	return orderItem
}

//向数据库添加订单详情
func AddOrderItem(orderItem *model.OrderItem) error {
	sqlStr := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		fmt.Println("AddOrderItem预编译和执行出错:", err)
		return err
	}
	return nil
}

//根据orderid获取订单详情
func GetOrderItems(orderId string) ([]*model.OrderItem, error) {
	//写sql语句
	sqlStr := "select count,amount,title,author,price,img_path from order_items where order_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, orderId)
	if err != nil {
		return nil, err
	}
	var orderItems []*model.OrderItem

	for rows.Next() {
		orderItem := &model.OrderItem{}
		err = rows.Scan(&orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, err
}

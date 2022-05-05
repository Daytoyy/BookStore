package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//写sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		fmt.Println("添加购物车预编译和执行出错:", err)
		return err
	}
	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		AddCartItem(cartItem)
	}

	return nil
}

//根据用户id获取购物车信息,每个用户都有独立的购物车
func GetCartByUserId(UserId string) (*model.Cart, error, int) {
	//写sql语句
	sqlStr := "select total_count,total_amount,id from carts where user_id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, UserId)

	cart := &model.Cart{}
	err := row.Scan(&cart.TotalCount, &cart.TotalAmount, &cart.CartId)

	if err != nil {
		return nil, err, 1
	}
	var cartId string
	cartId = cart.CartId
	//可根据cartid查询对应的购物项 cart.CartItems
	var cartItems []*model.CartItem
	cartItems, err = GetCartItemByCartId(cartId)
	if err != nil {
		return nil, err, 2
	}
	cart.CartItems = cartItems
	return cart, nil, 0
}

//更新购物车
func UpdateCart(cart *model.Cart) error {
	//写sql语句
	sqlStr := "update carts set total_count=?,total_amount=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.TotalCount, cart.TotalAmount, cart.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据购物车的id删除购物车
func DeleteCartByCartId(cartId string) error {
	//先删除该购物车id对应的购物项
	err := DelteCartItemsByCartId(cartId)
	sqlstr := "delete from carts where id=?"
	_, err = utils.Db.Exec(sqlstr, cartId)
	return err
}

//根据用户id获取购物车id
func GetCartIdByUserId(UserId int) (string, error) {
	//写sql语句
	sqlStr := "select id from carts where user_id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, UserId)

	var cartId string
	err := row.Scan(&cartId)
	fmt.Println("err=", err)
	return cartId, err
}

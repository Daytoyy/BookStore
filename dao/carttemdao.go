package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"strconv"
)

//向数据库中购物项的表中加入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//写sql语句
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.Id, cartItem.CartId)
	if err != nil {
		fmt.Println("添加购物项预编译和执行出错:", err)
		return err
	}
	return nil

}

//根据图书id获取对应的购物项
func GetCartItemByBookId(bookId string) (*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where book_id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookId)

	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.BookId, &cartItem.CartId)
	bookid := strconv.Itoa(cartItem.BookId)
	Book, _ := GetABook(bookid)
	cartItem.Book = Book

	if err != nil {
		return nil, err
	}
	return cartItem, err

}

//根据购物车id和图书获取购物车里bookid的购物项
func GetCartItemByCartIdAndBookId(cartId string, bookid string) ([]*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where cart_id=?&&book_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartId, bookid)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	var cartItem *model.CartItem
	for rows.Next() {
		cartItem = &model.CartItem{}
		rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.CartId)
		Book, _ := GetABook(bookid)
		cartItem.Book = Book
		//将cartItem添加到cartItems中
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, err
}

//根据购物车id和图书id获取购物车里的所有购物项
func GetCartItemByCartId(cartId string) ([]*model.CartItem, error) {
	//写sql语句
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartId)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	var cartItem *model.CartItem
	for rows.Next() {
		cartItem = &model.CartItem{}
		rows.Scan(&cartItem.CartItemId, &cartItem.Count, &cartItem.Amount, &cartItem.BookId, &cartItem.CartId)
		bookid := strconv.Itoa(cartItem.BookId)
		Book, _ := GetABook(bookid)
		cartItem.Book = Book
		//将cartItem添加到cartItems中
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, err
}

//根据用户id获取购物车里的购物项
func GetCartItemByUserId(UserId string) ([]*model.CartItem, error, int) {
	//先根据userid查找对应的CartId，再根据CartId获取购物车里的购物项
	//写sql语句
	sqlStr := "select id from carts where user_id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, UserId)
	var cartId string
	err := row.Scan(&cartId)
	if err != nil {
		return nil, err, 1
	}
	cartItem, err := GetCartItemByCartId(cartId)
	if err != nil {
		return nil, err, 2
	}
	return cartItem, nil, 0
}

//更新购物项
func UpdateCartItem(cartItem *model.CartItem) error {
	//写sql语句
	sqlStr := "update cart_items set count=?,amount=? where cart_id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.Amount, cartItem.CartId)
	if err != nil {
		return err
	}
	return nil
}

//根据bookid更新购物项
func UpdateCartItemByBookId(bookid int, cartItem *model.CartItem) error {
	//写sql语句
	sqlStr := "update cart_items set count=?,amount=? where book_id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.Amount, bookid)
	if err != nil {
		return err
	}
	return nil
}

//根据cartid从数据库删除cart_items
func DelteCartItemsByCartId(cartId string) error {
	sqlstr := "delete from cart_items where cart_id=?"
	_, err := utils.Db.Exec(sqlstr, cartId)
	return err
}

//根据cartitems的id从数据库中删除对应的cartitem项
func DelteCartItemById(cartitemId, userId string) error {
	sqlstr := "delete from cart_items where id=?"
	_, err := utils.Db.Exec(sqlstr, cartitemId)

	//更新购物车
	cartItems, _, _ := GetCartItemByUserId(userId)
	cart := &model.Cart{}
	var cnt int64
	var amt float64
	for _, v := range cartItems {
		cnt = cnt + v.Count
		amt = amt + v.Amount
	}
	cart.TotalCount = cnt
	cart.TotalAmount = amt
	//写sql语句
	sqlStr := "update carts set total_count=?,total_amount=? where user_id=?"
	//执行
	_, err = utils.Db.Exec(sqlStr, cart.TotalCount, cart.TotalAmount, userId)
	return err
}

//更新购物项
func UpdateCartItemByCount(bookId, cartId, bookCount string) (*model.Cart, float64, error) {

	//根据bookId查询购物项
	cartItem, _ := GetCartItemByBookId(bookId)

	count, _ := strconv.Atoi(bookCount)
	//fmt.Println("count=", count)
	amount := cartItem.Book.Price * float64(count)
	//fmt.Println("amount=", amount)
	cartItem.Count = int64(count)
	cartItem.Amount = model.FloatRound(amount, 2)

	//更新购物项
	bookid, _ := strconv.Atoi(bookId)
	err := UpdateCartItemByBookId(bookid, cartItem)
	//更新购物车
	cartItems, _ := GetCartItemByCartId(cartId)
	cart := &model.Cart{}
	var cnt int64
	var amt float64
	cnt = 0
	amt = 0.00
	for _, v := range cartItems {
		//fmt.Println("amt=", amt)
		//fmt.Println("v.Amount=", v.Amount)
		cnt = cnt + v.Count
		amt = amt + v.Amount
	}
	cart.TotalCount = cnt
	cart.TotalAmount = model.FloatRound(amt, 2)

	//fmt.Println("amt=", amt)
	//fmt.Println("cart.TotalAmount=", cart.TotalAmount)
	cart.CartId = cartId
	err = UpdateCart(cart)
	return cart, cartItem.Amount, err
}

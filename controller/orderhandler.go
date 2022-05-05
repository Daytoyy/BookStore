package controller

import (
	"bookstore/dao"
	"net/http"
	"text/template"
	// "strconv"
	// "go_code/goweb/bookstore/model"
)

//获取订单
func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	var userId int
	sess, flag := IsLogin(r)
	if flag == true {
		userId = sess.UserId
	}
	orders, _ := dao.GetOrders(userId)
	sess.Orders = orders

	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, sess)
}

//结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	var userId int
	sess, flag := IsLogin(r)
	if flag == true {
		userId = sess.UserId
	}
	//生成订单
	order := dao.CreatOrder(userId)
	//将订单数据加到数据库
	dao.AddOder(order)

	//根据获取到的购物车id更新图书销量和库存并生成订单详情
	cartItems, _ := dao.GetCartItemByCartId(order.CartId)
	for _, v := range cartItems {
		v.Book.Sales = v.Book.Sales + int(v.Count)
		v.Book.Stock = v.Book.Stock - int(v.Count)
		v.Book.Id = v.BookId
		dao.UpdateABook(v.Book)
		//生成订单详情
		orderItem := dao.CreatOrderItem(v)
		orderItem.OrderId = order.OrderId
		//将订单详情加到数据库
		dao.AddOrderItem(orderItem)
	}
	//清空该购物车
	dao.DeleteCartByCartId(order.CartId)
	//将orderId通过sess传给前端
	sess.OrderId = order.OrderId
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	t.Execute(w, sess)
}

// 获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	orderId := r.FormValue("orderId")
	OrderItems, _ := dao.GetOrderItems(orderId)
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	t.Execute(w, OrderItems)
}

//获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetAllOrders()
	var username string
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//获取session
		session, _ := dao.GetSession(cookieValue)
		username = session.UserName
	}

	if username == "admin" {
		t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
		t.Execute(w, orders)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/manager_err.html"))
		t.Execute(w, orders)
	}

}

//发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	//获取orderid
	orderId := r.FormValue("orderId")
	//设置订单状态state=1，更新到数据库
	state := "1"
	dao.SendOrderByOrderId(orderId, state)
	orders, _ := dao.GetAllOrders()
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	t.Execute(w, orders)
}

//处理订单
func TakeOrder(w http.ResponseWriter, r *http.Request) {

	var userId int
	sess, flag := IsLogin(r)
	if flag == true {
		userId = sess.UserId
	}
	//获取orderid和state
	orderId := r.FormValue("orderId")
	state := r.FormValue("state")
	//将订单状态更新到数据库
	dao.SendOrderByOrderId(orderId, state)

	orders, _ := dao.GetOrders(userId) //根据用户id获取当前用户的所有订单
	sess.Orders = orders

	//若订单取消 重新更新库存和销量
	if state == "3" {
		//根据orderItem的title获取book
		books, orderItems, _ := dao.GetABooksByOrderId(orderId)

		for t, v := range books {
			v.Sales = v.Sales - int(orderItems[t].Count)
			v.Stock = v.Stock + int(orderItems[t].Count)
			//v.Id=v.BookId
			dao.UpdateABook(v)
			//fmt.Println("book=", v)
		}
	}

	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	t.Execute(w, sess)
}

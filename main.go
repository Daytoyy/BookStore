package main

import (
	"net/http"

	"bookstore/controller"
)

func main() {
	//设置处理静态资源,如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))

	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	//首页
	//http.HandleFunc("/main",controller.IndexHandler)
	//http.HandleFunc("/main",IndexHandler)
	http.HandleFunc("/main", controller.QueryBooksByPrice)

	//登录
	http.HandleFunc("/login", controller.Login)

	//注册
	http.HandleFunc("/register", controller.Regist)

	//退出登录
	http.HandleFunc("/logout", controller.Logout)

	//重新切换账号登录
	http.HandleFunc("/relogin", controller.ReLogin)

	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/FindByUserName", controller.FindByUserName)

	//获取所有图书
	//http.HandleFunc("/getBooks",controller.GetBooksList)
	//分页后获取图书
	http.HandleFunc("/getPageBooks", controller.GetPageBooksList)

	//查询指定价格范围内的图书
	http.HandleFunc("/queryBooksByPrice", controller.QueryBooksByPrice)

	//通过书名或作者名查询出图书列表
	http.HandleFunc("/queryBooksByKeyWord", controller.QueryBooksByKeyword)

	//添加图书
	http.HandleFunc("/addBooks", controller.Addbooks)

	//删除图书
	http.HandleFunc("/deleteBook", controller.Deletebooks)

	//去更新图书的页面，在该页面获取要更新的图书的信息
	http.HandleFunc("/getAbookById", controller.GetAbookById)

	//更新图书
	http.HandleFunc("/updatebook", controller.Updatebook)

	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)

	//购物车
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)

	//删除购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)

	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)

	//结账
	http.HandleFunc("/checkout", controller.Checkout)

	//获取当前用户的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)

	//获取订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)

	//订单管理，获取所有用户订单
	http.HandleFunc("/getOrders", controller.GetOrders)

	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)

	//处理订单
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":8080", nil)
}

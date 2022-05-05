package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"encoding/json"
	"fmt"
	_ "fmt"
	"net/http"
	"strconv"
	"text/template"
)

//打开购物车页面
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	Cart := &model.Cart{}
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//获取session
		session, _ := dao.GetSession(cookieValue)
		if session.UserId > 0 {
			//根据登录的用户id从数据库获取图书
			userId := strconv.Itoa(session.UserId) //数据类型转换
			Cart, _, _ = dao.GetCartByUserId(userId)
			// Cart.UserId=session.UserId
			// Cart.UserName=session.UserName
			session.Cart = Cart
			t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
			t.Execute(w, session)
		} else {
			Cart.UserId = 0
		}
	}

}

//添加图书到购物车,将图书信息写入数据库
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {

	bookId := r.FormValue("bookId") //获取bookId
	book, _ := dao.GetABook(bookId) //根据bookId获取该图书信息
	sess, Islogin := IsLogin(r)     //判断是否处于登录状态
	if Islogin == true {
		Id := sess.UserId                              //获取当前登录状态下的用户id
		userId := strconv.Itoa(Id)                     //数据类型转换
		cart, err, flag := dao.GetCartByUserId(userId) //判断该用户状态下的数据库里是否能查到购物车
		if err != nil && flag != 1 {
			fmt.Println("查询数据库出错=", err)
			return
		}
		if flag == 1 { //获取不到购物车，创建购物车
			cartId := utils.CreateUUID() //生成uuid用作cartid
			cart := &model.Cart{
				CartId: cartId,
				UserId: Id,
			}
			var cartItems []*model.CartItem //创建一个购物项的切片用户购物车的购物项存放

			cartItem := &model.CartItem{ //创建一个购物项
				Book:   book, //将从前台获取到的图书信息存放进去
				Amount: book.Price,
				Count:  1,
				CartId: cartId,
			}
			cartItems = append(cartItems, cartItem) //将购物项添加到购物项切片
			cart.CartItems = cartItems
			cart.TotalCount = cartItem.Count
			cart.TotalAmount = cartItem.Amount
			dao.AddCart(cart) //写入数据库

		} else { //数据库中已有购物车
			//判断该用户的购物车是否已有该图书，根据查bookid判断
			ctms, _ := dao.GetCartItemByCartIdAndBookId(cart.CartId, bookId)
			if ctms == nil { //未查询到该图书的购物项
				cartItem := &model.CartItem{ //创建一个新的购物项添加这本图书
					Book:   book,
					Amount: book.Price,
					Count:  1,
					CartId: cart.CartId,
				}
				dao.AddCartItem(cartItem)               //将创建的购物项添加到数据库的购物项表中
				var cartItems []*model.CartItem         //创建一个购物项的切片用户购物车的购物项存放
				cartItems = append(cartItems, cartItem) //将购物项添加到购物项切片
				cart.CartItems = cartItems
				for _, v := range cart.CartItems {
					cart.TotalCount = cart.TotalCount + v.Count
					cart.TotalAmount = cart.TotalAmount + v.GetAmount()
				}

			} else { //查询到该购物车已有该图书的购物项
				for _, v := range cart.CartItems {
					if v.BookId == book.Id { //在数据库里找到与要添加的图书id相同的购物项
						v.Count = v.Count + 1                    //该购物项的数量小计
						v.Amount = float64(v.Count) * book.Price //该购物项的金额小计
						dao.UpdateCartItemByBookId(v.BookId, v)  //更新购物项
					}
				}
				var cnt int64
				var amt float64
				for _, v := range cart.CartItems {
					cnt = cnt + v.Count  //购物车里的总数量
					amt = amt + v.Amount //购物车里的总金额
				}
				cart.TotalCount = cnt
				cart.TotalAmount = model.FloatRound(amt, 2)
				cnt = 0
				amt = 0
			}
			dao.UpdateCart(cart) //更新图书
		}
		w.Write([]byte(book.Title))
	} else {
		w.Write([]byte("请先进行登录再操作！"))
	}

}

//清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	cartId := r.FormValue("cartId")
	dao.DeleteCartByCartId(cartId)
	GetCartInfo(w, r)
}

//删除该购物车的购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemId := r.FormValue("cartItemId")

	sess, flag := IsLogin(r)
	if flag == true {
		//根据登录的用户id从数据库获取图书
		userId := strconv.Itoa(sess.UserId) //数据类型转换
		dao.DelteCartItemById(cartItemId, userId)
	}
	GetCartInfo(w, r)
}

// 更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取cartItemId
	bookId := r.FormValue("bookId")
	//获取设置的bookCount
	bookCount := r.FormValue("bookCount")
	cartId := r.FormValue("cartId")
	cart, amount, _ := dao.UpdateCartItemByCount(bookId, cartId, bookCount)

	data := model.Data{
		TotalAmount: cart.TotalAmount,
		TotalCount:  cart.TotalCount,
		Amount:      amount,
	}

	json, _ := json.Marshal(data)
	w.Write(json)
}

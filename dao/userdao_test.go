package dao

import (
	"bookstore/model"
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	// fmt.Println("测试userdao中的函数")
	// t.Run("验证用户名或密码：",testLogin)
	// t.Run("验证用户名：",testRegist)
	// t.Run("保存用户：",testSave)
	//t.Run("测试添加图书",testAddBooks)

	// t.Run("测试删除图书",testDelbook)
	//t.Run("测试根据id获取一本图书",testGetABook)
	//t.Run("测试根据id更新图书",testUpdateBook)
	//t.Run("测试获取数据库中的所有图书",testGetBooks)
	//t.Run("测试获取分页后的图书信息",testGetPageBooks)
	//t.Run("测试获取一个价格区间内的数据库里的分页后的图书信息",testGetPageBooksByPrice)
	//t.Run("测试添加sesion",testAddsession)
	//t.Run("测试通过SessionId删除sesion",testDelsession)
	//t.Run("测试通过sessionid获取数据库中的session数据",testGetSession)

	//t.Run("测试添加图书到购物车",testAddCarts)
	//t.Run("测试根据图书id获取对应的购物项",testGetCartItemsByBookId)
	//t.Run("测试根据购物车id获取购物车里所有的购物项",testGetCartItemsByCartId)
	//t.Run("根据用户id获取对应购物车里的所有项",testGetCartItemsByUserId)
	//t.Run("测试根据购物车id获取购物车里所有的购物项", testGetCartItemsByCartIdAndBookId)
}

func testLogin(t *testing.T) {
	user, _ := CheckUsernameAndPssword("admin", "admin")
	fmt.Println("获取的用户信息是：", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin2")
	fmt.Println("获取的用户信息是：", user)
}

func testSave(t *testing.T) {
	SaveUser("admin3", "66666", "wym@666.com")

}

//测试获取图书
func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本图书信息是：%v\n", k+1, v)
	}
}

//测试添加图书
func testAddBooks(t *testing.T) {
	book := &model.Book{
		Title:   "Go语言入门经典",
		Author:  "张海燕",
		Price:   59.00,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	Addbook(book)

}

//测试删除图书
func testDelbook(t *testing.T) {
	Delbook("43")
}

//测试根据id获取一本图书信息
func testGetABook(t *testing.T) {
	book, _ := GetABook("1")

	fmt.Println(book)

}

//测试更新图书
func testUpdateBook(t *testing.T) {
	book := &model.Book{
		Id:      40,
		Title:   "Go语言学习笔记1.0",
		Author:  "雨痕",
		Price:   59.99,
		Sales:   100,
		Stock:   100,
		ImgPath: "/static/img/default.jpg",
	}
	UpdateABook(book)

}

//测试获取分页后的图书信息
func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("1")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数为：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	for t, v := range page.Books {
		fmt.Printf("当前页的第%d条图书信息：%v\n", t+1, v)
	}
}

//测试获取一个价格区间内的数据库里的分页后的图书信息
func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("60", "70", "1")
	fmt.Println("当前页是：", page.PageNo)
	fmt.Println("总页数为：", page.TotalPageNo)
	fmt.Println("总记录数是：", page.TotalRecord)
	for t, v := range page.Books {
		fmt.Printf("当前页的第%d条图书信息：%v\n", t+1, v)
	}
}

//测试添加sesion
func testAddsession(t *testing.T) {
	sess := &model.Session{
		SessionId: "54312123456",
		UserName:  "admin",
		UserId:    15,
	}
	AddSession(sess)

}

//测试测试通过SessionId删除sesion
func testDelsession(t *testing.T) {
	DelSession("6a66ba3e-e6fd-4fa4-55a8-dbe8322b813d")
}

//测试通过sessionid获取数据库中的session数据
func testGetSession(t *testing.T) {
	sess, _ := GetSession("54abe639-43fc-4523-4eef-147f0702a2d8")
	fmt.Println(sess)
}

//测试添加图书到购物车
func testAddCarts(t *testing.T) {

	book1 := &model.Book{
		Id:    5,
		Price: 19,
	}
	book2 := &model.Book{
		Id:    6,
		Price: 29,
	}
	//创建购物项
	var cartItems []*model.CartItem
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  10,
		CartId: "12312345",
	}
	cartItems = append(cartItems, cartItem1)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartId: "12312345",
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartId:    "12312345",
		CartItems: cartItems,
		UserId:    2,
	}
	err := AddCart(cart)
	if err != nil {
		return
	}
	fmt.Println("成功添加图书")
}

//测试根据图书id获取对应的购物项
func testGetCartItemsByBookId(t *testing.T) {
	cartItem, err := GetCartItemByBookId("1")
	fmt.Println("err", err)
	fmt.Println("该bookid对应的购物项是：", cartItem)
}

//测试根据购物车id获取购物车里所有的购物项
// func testGetCartItemsByCartId(t *testing.T){
// 	cartItems,err:=GetCartItemByCartId("87654321")
// 	println("err:",err)
// 	for k,v:=range cartItems{
// 		fmt.Printf("该购物车里的第%d个购物项是:%v\n",k+1,v)
// 	}
// }
//根据用户id获取对应购物车里的所有项
func testGetCartItemsByUserId(t *testing.T) {
	cartItems, err, flag := GetCartItemByUserId("6")
	if err != nil {
		println("err:", err)
	}
	println("falg:", flag)
	for k, v := range cartItems {
		fmt.Printf("该用户对应购物车里的第%d个购物项是:%v\n", k+1, v)
	}
}

// func testGetCartItemsByCartIdAndBookId(t *testing.T) {
// 	cartItems, err := GetCartItemByCartId("460b83ec-525c-411a-6afd-c1ed299f0bb0", "12")
// 	if err != nil {
// 		println("err:", err)
// 	}
// 	for k, v := range cartItems {
// 		fmt.Printf("该用户对应购物车里的第%d个购物项是:%v\n", k+1, v)
// 	}
// }

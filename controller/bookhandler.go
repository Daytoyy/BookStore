package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"net/http"
	"strconv"
	"text/template"
)

//显示首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	//获取页码
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)

	//解析模板,使用Must函数处理异常
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)

}

//获取所有图书信息
func GetBooksList(w http.ResponseWriter, r *http.Request) {
	//调用bookdao获取所有图书
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}

//获取分页后的图书信息
func GetPageBooksList(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageBooks(pageNo)

	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//获取session
		session, _ := dao.GetSession(cookieValue)
		if session.UserId > 0 {
			page.IsLogin = true
			page.UserName = session.UserName
		} else {
			page.IsLogin = false
		}
	}
	if page.UserName == "admin" {
		t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
		t.Execute(w, page)
	} else {
		t := template.Must(template.ParseFiles("views/pages/manager/manager_err.html"))
		t.Execute(w, page)
	}

}

//添加图书 Addbook
func Addbooks(w http.ResponseWriter, r *http.Request) {
	//从提交的表单中获取要添加的图书信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//做进制转换
	fprice, _ := strconv.ParseFloat(price, 64)
	fsales, _ := strconv.ParseInt(sales, 10, 0)
	fstock, _ := strconv.ParseInt(stock, 10, 0)
	//创建book
	book := &model.Book{
		Title:   title,
		Author:  author,
		Price:   fprice,
		Sales:   int(fsales),
		Stock:   int(fstock),
		ImgPath: "/static/img/default.jpg",
	}
	dao.Addbook(book)

	GetPageBooksList(w, r)
}

//删除图书
func Deletebooks(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	dao.Delbook(bookId)
	GetPageBooksList(w, r)
}

//更新图书信息
func Updatebook(w http.ResponseWriter, r *http.Request) {
	//从提交的表单中获取图书信息
	bookId := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//做进制转换
	fprice, _ := strconv.ParseFloat(price, 64)
	fsales, _ := strconv.ParseInt(sales, 10, 0)
	fstock, _ := strconv.ParseInt(stock, 10, 0)
	fbookId, _ := strconv.ParseInt(bookId, 10, 0)
	//创建book
	book := &model.Book{
		Id:      int(fbookId),
		Title:   title,
		Author:  author,
		Price:   fprice,
		Sales:   int(fsales),
		Stock:   int(fstock),
		ImgPath: "/static/img/default.jpg",
	}
	dao.UpdateABook(book)

	GetPageBooksList(w, r)
}

//通过id获取一本图书信息
func GetAbookById(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	book, _ := dao.GetABook(bookId)
	t := template.Must(template.ParseFiles("views/pages/manager/book_modify.html"))
	t.Execute(w, book)

}

//获取带分页和价格范围的图书
func QueryBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//从提交的表单中获取要查询的图书价格范围
	Min := r.FormValue("min")
	Max := r.FormValue("max")
	//获取页码
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if Min == "" || Max == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		page, _ = dao.GetPageBooksByPrice(Min, Max, pageNo)
		page.MinPrice = Min
		page.MaxPrice = Max
	}
	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//获取session
		session, _ := dao.GetSession(cookieValue)
		if session.UserId > 0 {
			page.IsLogin = true
			page.UserName = session.UserName
		} else {
			page.IsLogin = false
		}
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

// 通过输入的关键字获取图书信息
func QueryBooksByKeyword(w http.ResponseWriter, r *http.Request) {
	keyword := r.FormValue("keyword")
	option := r.FormValue("optionlist")
	//获取页码
	pageNo := r.FormValue("PageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	// fmt.Println(option, keyword)
	var page *model.Page

	if keyword == "" || option == "" {
		page, _ = dao.GetPageBooks(pageNo)
	} else {
		if option == "bookAuthor" {
			page, _ = dao.GetBookByNameAndAuthor(1, keyword, pageNo)
		} else {
			page, _ = dao.GetBookByNameAndAuthor(0, keyword, pageNo)
		}
	}

	//获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		//获取session
		session, _ := dao.GetSession(cookieValue)
		if session.UserId > 0 {
			page.IsLogin = true
			page.UserName = session.UserName
		} else {
			page.IsLogin = false
		}
	}

	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, page)
}

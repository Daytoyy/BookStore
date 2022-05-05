package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"strconv"
)

//操作数据库

//从数据库获取图书信息
func GetBooks() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}
	return books, err
}

//往数据库添加图书
func Addbook(b *model.Book) error {
	//写sql语句
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//根据Id从数据库删除图书
func Delbook(id string) error {
	//写sql语句
	sqlStr := "delete from books where id=?"
	_, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return err
	}

	return nil
}

//从数据库获取一本图书信息
func GetABook(id string) (*model.Book, error) {
	//写sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, id) //QueryRow执行一次返回一行结果
	book := &model.Book{}
	//给book中的字段赋值
	row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

//更新图书
func UpdateABook(book *model.Book) error {
	//写sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=?,img_path=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath, book.Id)
	if err != nil {
		return err
	}
	return nil
}

//获取分页后的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	//将页码转成int64
	IpageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//先获取数据库中的总记录数
	sqlStr := "select count(*) from books"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64
	pageSize = 4
	//获取总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//获取当前页中的图书
	sqlStr = "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr, (IpageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)

	}

	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      IpageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func GetPageBooksByPrice(minPrice, maxPrice, pageNo string) (*model.Page, error) {
	//将页码转成int64
	IpageNo, _ := strconv.ParseInt(pageNo, 10, 64)

	//先获取数据库中的总记录数
	sqlStr := "select count(*) from books where price between ? and ?"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64
	pageSize = 4
	//获取总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//获取当前页中的图书
	sqlStr = "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr, minPrice, maxPrice, (IpageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)

	}
	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      IpageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//根据orderId获取books
func GetABooksByOrderId(orderId string) ([]*model.Book, []*model.OrderItem, error) {
	//先从order_items获取title
	//写sql语句
	sqlStr := "select count,title from order_items where order_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, orderId)
	var books []*model.Book
	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.Count, &orderItem.Title)
		orderItems = append(orderItems, orderItem)

	}

	for _, v := range orderItems {
		//写sql语句
		sqlStr := "select id,title,author,price,sales,stock,img_path from books where title=?"
		//执行
		rows, err = utils.Db.Query(sqlStr, v.Title)

		for rows.Next() {
			book := &model.Book{}
			//给book中的字段赋值
			rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
			//将book添加到books中
			books = append(books, book)
		}
	}

	return books, orderItems, err
}

//从数据库获取所有图书的书名和作者信息
func GetBookNameAndAuthor() ([]*model.Book, error) {
	//写sql语句
	sqlStr := "select title,author from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		//给book中的字段赋值
		rows.Scan(&book.Title, &book.Author)
		//将book添加到books中
		books = append(books, book)
	}
	return books, err
}

//根据作者名或书名从数据库获取符合条件的所有图书信息
func GetBookByNameAndAuthor(key int, temp string, pageNo string) (*model.Page, error) {
	//将页码转成int64
	IpageNo, _ := strconv.ParseInt(pageNo, 10, 64)

	var sqlStr string

	//设置一个变量接收总记录数 设置每页只显示4条记录
	var totalRecord, pageSize int64
	pageSize = 4

	if key == 1 {
		//先获取数据库中的总记录数
		sqlStr = "select count(*) from books where author=?"
		//执行
		row := utils.Db.QueryRow(sqlStr, temp)
		row.Scan(&totalRecord)
		//写sql语句
		sqlStr = "select id,title,author,price,sales,stock,img_path from books where author=?"
	} else {

		//先获取数据库中的总记录数
		sqlStr = "select count(*) from books where title=?"
		//执行
		row := utils.Db.QueryRow(sqlStr, temp)
		row.Scan(&totalRecord)

		//写sql语句
		sqlStr = "select id,title,author,price,sales,stock,img_path from books where title=?"
		//执行

	}
	rows, err := utils.Db.Query(sqlStr, temp)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)

	}
	//获取总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}
	//创建page
	page := &model.Page{
		Books:       books,
		PageNo:      IpageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}

	return page, err
}

package model

type Page struct {
	Books       []*Book //存放每页查询出来的图书
	PageNo      int64   //当前页
	PageSize    int64   //每页显示的条数
	TotalPageNo int64   //总页数，通过计算得出
	TotalRecord int64   //总记录数，通过查询数据库得到
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	UserName    string
}

func (page *Page) IsHasPrev() bool {
	return page.PageNo > 1
}

func (page *Page) IsHasNext() bool {
	return page.PageNo < page.TotalPageNo
}

func (page *Page) GetPrevPageNo() int64 {
	if page.IsHasPrev() {
		return page.PageNo - 1
	} else {
		return 1
	}
}

func (page *Page) GetNextPageNo() int64 {
	if page.IsHasNext() {
		return page.PageNo + 1
	} else {
		return page.PageNo
	}
}

func (page *Page) IndexPage() int64 {
	return page.PageNo
}

func (page *Page) IsAdmin() bool {
	if page.UserName == "admin" {
		return true
	} else {
		return false
	}

}

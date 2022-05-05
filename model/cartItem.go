package model

//购物项结构体，每个购物车中关于物图书的各项描述
type CartItem struct {
	CartItemId int64   //购物项的id(排列序号,系统自动分配)
	Book       *Book   //图书信息
	Count      int64   //图书数量
	Amount     float64 //购物项中图书的金额小计，通过计算得到,等于该图书的数量x该图书的价格
	CartId     string  //当前购物项属于哪一个购物车 用户id
	BookId     int
}

//获取购物车里同一种图书的价格小计 该图书的数量x该图书的价格 Amount
func (cartItem *CartItem) GetAmount() float64 {
	Icount := float64(cartItem.Count)
	return Icount * cartItem.Book.Price
}

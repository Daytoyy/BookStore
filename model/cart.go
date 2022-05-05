package model

//购物车 每个用户登录后有自己单独的购物车
type Cart struct {
	CartId      string      //购物车的id
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物车中图书的总数量
	TotalAmount float64     //购物车中图书的总金额
	UserId      int         //当前购物车所属的用户
	UserName    string      //当前购物车所属的用户名
}

//获取购物车中图书的总数量
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount
}

//获取购物车中的总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}

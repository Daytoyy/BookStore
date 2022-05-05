package model

//定义sessiom结构体
type Session struct {
	SessionId string
	UserName  string
	UserId    int
	Cart      *Cart
	OrderId   string
	Orders    []*Order
}

//判断是否是管理员用户
func (session *Session) IsAdmin() bool {
	if session.UserName == "admin" {
		return true
	} else {
		return false
	}

}

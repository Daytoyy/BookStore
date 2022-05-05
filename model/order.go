package model

type Order struct {
	OrderId     string
	CreateTime  string //订单生成时间
	TotalCount  int64
	TotalAmount float64
	State       int64 //0-未发货 1-已发货 2-交易完成 3-订单取消
	UserId      int64
	UserName    string
	CartId      string
}

func (order *Order) NoSend() bool {
	if order.State == 0 {
		return true
	} else {
		return false
	}
}

func (order *Order) SendComplate() bool {
	if order.State == 1 {
		return true
	} else {
		return false
	}
}

func (order *Order) Complate() bool {
	if order.State == 2 {
		return true
	} else {
		return false
	}
}

func (order *Order) Cancel() bool {
	if order.State == 3 {
		return true
	} else {
		return false
	}
}

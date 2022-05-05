package model

import(
	"strconv"
	"fmt"
)

type Data struct{
	TotalAmount float64
	TotalCount int64
	Amount float64
}

//处理数据精度问题
func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}
package main

import (
	"github.com/os369510/taiextools/twse"
)

//var source = map[string]int{
//	twse: 1,
//}

func main() {
	var api string = "STOCK_DAY"
	var date string = "date=20180711"
	var stockNo string = "stockNo=9958"

	err := twse.TwseQuery(api, date, stockNo)
	if err != nil {
		panic(err)
	}
}

package twse

import (
	"fmt"
)

var base string = "http://www.twse.com.tw/exchangeReport"
var resp_json string = "response=json"

type TwseBaseObj struct {
	source  string
	handler TwseHandler
}

type TwseHandler interface {
	query() (err error)
	show() (err error)
}

func TwseQuery(api string, date string, stockNo string) (err error) {
	var tbo TwseBaseObj
	switch api {
	case "BWIBBU":
		tbo.handler, err = InitBWIBBUObj(base, resp_json, date, stockNo)
	case "STOCK_DAY":
		tbo.handler, err = InitStockDayObj(base, resp_json, date, stockNo)
	default:
		err = fmt.Errorf("Unkonw ", api)
	}
	if err != nil {
		return
	}

	err = tbo.handler.query()
	if err != nil {
		return
	}

	err = tbo.handler.show()
	if err != nil {
		return
	}
	return nil
}

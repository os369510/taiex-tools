package twse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var stock_day_api string = "STOCK_DAY"

type StockDayReq struct {
	url      string
	api      string
	base     string
	response string
	date     string
	stockNo  string
}

type StockDayResp struct {
	Stat   string     `json:"stat"`
	Date   string     `json:"date"`
	Title  string     `json:"title"`
	Fields []string   `json:"fields"`
	Data   [][]string `json:"data"`
	Notes  []string   `json:"notes"`
}

type StockDayObj struct {
	req  StockDayReq
	resp StockDayResp
}

func (obj *StockDayObj) query() (err error) {
	var myClient = &http.Client{Timeout: 10 * time.Second}

	obj.req.url = obj.req.base + "/" + obj.req.api + "?" + obj.req.response + "&" + obj.req.date + "&" + obj.req.stockNo
	fmt.Println("Request: \"" + obj.req.url + "\"")

	url_resp, err := myClient.Get(obj.req.url)
	if err != nil {
		return
	}
	defer url_resp.Body.Close()

	err = json.NewDecoder(url_resp.Body).Decode(&obj.resp)
	if err != nil {
		return
	}

	return
}

func (obj *StockDayObj) show() (err error) {
	if obj.resp.Stat != "OK" {
		err = fmt.Errorf("The stat of response is %s not \"OK\".\n", obj.resp.Stat)
		return
	}
	fmt.Printf("Date: \t%s\n", obj.resp.Date)
	fmt.Printf("Title: \t%s\n", obj.resp.Title)
	for i, field := range obj.resp.Fields {
		if i != 0 {
			fmt.Printf("\t")
		}
		fmt.Printf("%s", field)
	}
	fmt.Printf("\n")
	for _, row := range obj.resp.Data {
		for j, field := range row {
			if j != 0 {
				fmt.Printf("\t")
			}
			fmt.Printf("%s", field)
		}
		fmt.Printf("\n")
	}
	return nil
}

func InitStockDayObj(base string, response string, date string, stockNo string) (obj *StockDayObj, err error) {
	obj = &StockDayObj{}
	obj.req.api = stock_day_api
	obj.req.base = base
	obj.req.response = response
	obj.req.date = date
	obj.req.stockNo = stockNo
	return
}

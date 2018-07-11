package twse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var bwibbu_api string = "BWIBBU"

type BWIBBUReq struct {
	url      string
	api      string
	base     string
	response string
	date     string
	stockNo  string
}

type BWIBBUResp struct {
	Stat   string           `json:"stat"`
	Date   string           `json:"date"`
	Title  string           `json:"title"`
	Fields []string         `json:"fields"`
	Data   []BWIBBURespData `json:"data"`
	Notes  []string         `json:"notes"`
}

type BWIBBURespData struct {
	Date  string // Date
	Dy    string // Dividend yield
	DYear uint32 // Dividend Year
	PE    string // Price-Earnings Ratio
	PBR   string // Price-Book Ratio
	FRYS  string // Financial Report (Year/Season)
}

type BWIBBUObj struct {
	req  BWIBBUReq
	resp BWIBBUResp
}

func (obj *BWIBBUObj) query() (err error) {
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

func (obj *BWIBBUObj) show() (err error) {
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
		fmt.Printf("%s\t%s\t%d\t%s\t%s\t%s\n", row.Date, row.Dy, row.DYear, row.PE, row.PBR, row.FRYS)
	}
	return nil
}

func InitBWIBBUObj(base string, response string, date string, stockNo string) (obj *BWIBBUObj, err error) {
	obj = &BWIBBUObj{}
	obj.req.api = bwibbu_api
	obj.req.base = base
	obj.req.response = response
	obj.req.date = date
	obj.req.stockNo = stockNo
	return
}

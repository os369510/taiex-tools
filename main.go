package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TWSE_Request struct {
	url      string
	response string
	date     string
	stockNo  string
}

type TWSE_Response struct {
	stat   string     `json:"stat,omitempty"`
	date   string     `json:"date,omitempty"`
	title  string     `json:"title,omitempty"`
	fields []string   `json:"fields,omitempty"`
	data   [][]string `json:"data,omitempty"`
	notes  []string   `json:"notes,omitempty"`
}

var twse_url string = "http://www.twse.com.tw/exchangeReport/STOCK_DAY?"
var twse_response string = "response=json"

var myClient = &http.Client{Timeout: 10 * time.Second}

func main() {
	var twse_req TWSE_Request
	twse_req.url = twse_url
	twse_req.response = twse_response
	twse_req.date = "date=20180711"
	twse_req.stockNo = "stockNo=9958"

	resp, err := myClient.Get(twse_req.url + "&" + twse_req.response + "&" + twse_req.date + "&" + twse_req.stockNo)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	twse_resp := new(TWSE_Response)

	err = json.NewDecoder(resp.Body).Decode(&twse_resp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", twse_resp)
}

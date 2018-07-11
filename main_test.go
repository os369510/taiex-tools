package main

import (
	"github.com/os369510/taiextools/twse"
	"testing"
)

func TestSTOCK_DAY(t *testing.T) {
	var date string = "date=20180710"
	var api string = "STOCK_DAY"
	var stockNo string = "stockNo=9958"

	err := twse.TwseQuery(api, date, stockNo)
	if err != nil {
		t.Error(err)
	}
}
func TestBWIBBU(t *testing.T) {
	var date string = "date=20180710"
	var api string = "BWIBBU"
	var stockNo string = "stockNo=9958"

	err := twse.TwseQuery(api, date, stockNo)
	if err != nil {
		t.Error(err)
	}
}

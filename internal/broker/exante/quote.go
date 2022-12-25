package exante

import (
	"fmt"
	"strconv"
)

type Quote struct {
	Timestamp int64  `json:"timestamp"`
	SymbolID  string `json:"symbolId"`
	Event     string `json:"event"`
	Bid       []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"bid"`
	Ask []struct {
		Price string `json:"price"`
		Size  string `json:"size"`
	} `json:"ask"`
}

func (q Quote) Bid_() (price, size float64, err error) {
	if len(q.Bid) == 0 {
		err = fmt.Errorf("bid must be defined")
		return
	}
	val, err1 := strconv.ParseFloat(q.Bid[0].Price, 64)
	if err1 != nil {
		err = fmt.Errorf("cannot parse bid price: %w", err1)
		return
	}
	price = val
	val, err1 = strconv.ParseFloat(q.Bid[0].Size, 64)
	if err1 != nil {
		err = fmt.Errorf("cannot parse bid size: %w", err1)
		return
	}
	size = val
	return
}

func (q Quote) Ask_() (price, size float64, err error) {
	if len(q.Ask) == 0 {
		err = fmt.Errorf("ask must be defined")
		return
	}
	val, err1 := strconv.ParseFloat(q.Ask[0].Price, 64)
	if err1 != nil {
		err = fmt.Errorf("cannot parse ask price: %w", err1)
		return
	}
	price = val
	val, err1 = strconv.ParseFloat(q.Ask[0].Size, 64)
	if err1 != nil {
		err = fmt.Errorf("cannot parse ask size: %w", err1)
		return
	}
	size = val
	return
}

func (q Quote) Symbol_() string {
	return q.SymbolID
}

func (q Quote) Timestamp_() int64 {
	return q.Timestamp
}

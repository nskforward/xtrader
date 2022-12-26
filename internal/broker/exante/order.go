package exante

/*
{
  "accountId": "ABC1234.001",
  "symbolId": "AAPL.NASDAQ",
  "side": "buy",
  "quantity": "6",
  "orderType": "market",
  "stopPrice": "100.0",
  "limitPrice": "120.0",
  "partQuantity": "1.0",
  "placeInterval": "1",
  "priceDistance": "1",
  "duration": "day",
  "gttExpiration": "2017-08-14T02:40:00.000Z",
  "clientTag": "some client tag",
  "takeProfit": "130.0",
  "stopLoss": "90.0",
  "ocoGroup": "d440b5b6-a40f-44e5-8c3b-a9a419fea7b3",
  "ifDoneParentId": "3a5bf47e-ec54-4782-b4e3-0091164c7c71"
}
*/

type Order struct {
	AccountID      string        `json:"accountId"`
	ClientTag      string        `json:"clientTag,omitempty"`
	TakeProfit     string        `json:"takeProfit,omitempty"`
	StopLoss       string        `json:"stopLoss,omitempty"`
	SymbolID       string        `json:"symbolId"`
	Side           OrderSide     `json:"side"`
	Quantity       string        `json:"quantity"`
	OrderType      OrderType     `json:"orderType"`
	Duration       OrderDuration `json:"duration"`
	StopPrice      string        `json:"stopPrice,omitempty"`
	LimitPrice     string        `json:"limitPrice,omitempty"`
	PartQuantity   string        `json:"partQuantity,omitempty"`
	PlaceInterval  string        `json:"placeInterval,omitempty"`
	PriceDistance  string        `json:"priceDistance,omitempty"`
	GttExpiration  string        `json:"gttExpiration,omitempty"`
	OcoGroup       string        `json:"ocoGroup,omitempty"`
	IfDoneParentId string        `json:"ifDoneParentId,omitempty"`
}

type OrderParameters struct {
	SymbolID       string        `json:"symbolId"`
	Side           OrderSide     `json:"side"`
	Quantity       string        `json:"quantity"`
	OrderType      OrderType     `json:"orderType"`
	Duration       OrderDuration `json:"duration"`
	StopPrice      string        `json:"stopPrice,omitempty"`
	LimitPrice     string        `json:"limitPrice,omitempty"`
	PartQuantity   string        `json:"partQuantity,omitempty"`
	PlaceInterval  string        `json:"placeInterval,omitempty"`
	PriceDistance  string        `json:"priceDistance,omitempty"`
	GttExpiration  string        `json:"gttExpiration,omitempty"`
	OcoGroup       string        `json:"ocoGroup,omitempty"`
	IfDoneParentId string        `json:"ifDoneParentId,omitempty"`
}

type OrderSide string

const (
	Buy  OrderSide = "buy"
	Sell OrderSide = "sell"
)

type OrderType string

const (
	Market       OrderType = "market"
	Limit        OrderType = "limit"
	Stop         OrderType = "stop"
	StopLimit    OrderType = "stop_limit"
	Twap         OrderType = "twap"
	TrailingStop OrderType = "trailing_stop"
	Iceberg      OrderType = "iceberg"
)

type OrderDuration string

const (
	Day               OrderDuration = "day"
	FillOrKill        OrderDuration = "fill_or_kill"
	ImmediateOrCancel OrderDuration = "immediate_or_cancel"
	GoodTillCancel    OrderDuration = "good_till_cancel"
	GoodTillTime      OrderDuration = "good_till_time"
	AtTheOpening      OrderDuration = "at_the_opening"
	AtTheClose        OrderDuration = "at_the_close"
)

package exante

type OrderResponse struct {
	OrderID               string          `json:"orderId"`
	PlaceTime             string          `json:"placeTime"`
	CurrentModificationId string          `json:"currentModificationId"`
	AccountID             string          `json:"accountId"`
	Username              string          `json:"username"`
	ClientTag             string          `json:"clientTag"`
	OrderState            OrderState      `json:"orderState"`
	OrderParameters       OrderParameters `json:"orderParameters"`
}

type OrderState struct {
	Status     string `json:"status"`
	LastUpdate string `json:"lastUpdate"`
	Reason     string `json:"reason"`
	Fills      []struct {
		Timestamp string `json:"timestamp"`
		Quantity  string `json:"quantity"`
		Price     string `json:"price"`
		Position  string `json:"position"`
	} `json:"fills"`
}

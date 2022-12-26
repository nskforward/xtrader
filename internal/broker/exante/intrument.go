package exante

/*
{
    "optionData": {
      "optionGroupId": "INTC.CBOE.6F2023.P*",
      "strikePrice": "30.5",
      "optionRight": "PUT"
    },
    "name": "Intel",
    "symbolId": "INTC.CBOE.6F2023.P30_5",
    "description": "Option on Intel 6 Jan 2023 PUT 30.5",
    "icon": null,
    "underlyingSymbolId": "INTC.NASDAQ",
    "country": "US",
    "identifiers": null,
    "exchange": "CBOE",
    "symbolType": "OPTION",
    "currency": "USD",
    "minPriceIncrement": "0.01",
    "ticker": "INTC",
    "expiration": 1673038800000,
    "group": "INTC"
  }
*/

type Instrument struct {
	Name               string `json:"name"`
	SymbolID           string `json:"symbolId"`
	Description        string `json:"description"`
	SymbolType         string `json:"symbolType"`
	Expiration         int64  `json:"expiration"`
	Ticker             string `json:"ticker"`
	Currency           string `json:"currency"`
	UnderlyingSymbolId string `json:"underlyingSymbolId"`
	Group              string `json:"group"`
	MinPriceIncrement  string `json:"minPriceIncrement"`
	Exchange           string `json:"exchange"`
	Country            string `json:"country"`
	OptionData         struct {
		OptionGroupId string `json:"optionGroupId"`
		StrikePrice   string `json:"strikePrice"`
		OptionRight   string `json:"optionRight"`
	} `json:"optionData"`
}

package broker

type Quote interface {
	Bid_() (price, size float64, err error)
	Ask_() (price, size float64, err error)
	Symbol_() string
	Timestamp_() int64
}

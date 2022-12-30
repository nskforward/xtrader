package broker

type Config struct {
	Name   string `json:"name"`
	Exante struct {
		Addr      string `json:"addr"`
		AccountID string `json:"account_id"`
		ClientID  string `json:"client_id"`
		AppID     string `json:"app_id"`
		SharedKey string `json:"shared_key"`
	} `json:"exante"`
}

package exante

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetActiveOrders(ctx context.Context, cfg Config, symbol string) error {
	url := fmt.Sprintf("%s/trade/3.0/orders/active?symbolId=%s", cfg.Addr, symbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.WithContext(ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", NewToken(cfg)}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode > 399 {
		return fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	fmt.Println(string(data))

	return nil
}

package exante

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func PlaceOrder(ctx context.Context, cfg Config, order Order) (orderResponse []OrderResponse, err error) {
	url := fmt.Sprintf("%s/trade/3.0/orders", cfg.Addr)
	dataReq, err2 := json.Marshal(order)
	if err2 != nil {
		err = err2
		return
	}
	req, err2 := http.NewRequest("POST", url, bytes.NewReader(dataReq))
	if err2 != nil {
		err = err2
		return
	}
	req.WithContext(ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", NewToken(cfg)}, " "))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(dataReq)))

	resp, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		err = err2
		return
	}

	defer resp.Body.Close()

	dataResp, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
		return
	}

	if resp.StatusCode > 399 {
		err = fmt.Errorf("bad http response code: %s: %s", resp.Status, string(dataResp))
		return
	}

	err = json.Unmarshal(dataResp, &orderResponse)
	return
}

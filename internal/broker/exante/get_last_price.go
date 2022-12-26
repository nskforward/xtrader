package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetLastPrice(ctx context.Context, cfg Config, symbol string) (quote Quote, err error) {
	url := fmt.Sprintf("%s/md/3.0/feed/%s/last", cfg.Addr, symbol)
	req, err2 := http.NewRequest("GET", url, nil)
	if err2 != nil {
		err = err2
		return
	}
	req.WithContext(ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", NewToken(cfg)}, " "))
	req.Header.Add("Accept", "application/json")
	resp, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		err = err2
		return
	}
	defer resp.Body.Close()

	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
		return
	}

	if resp.StatusCode > 399 {
		err = fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
		return
	}

	var qs []Quote

	err2 = json.Unmarshal(data, &qs)
	if err2 != nil {
		err = err2
		return
	}

	if len(qs) == 0 {
		err = fmt.Errorf("no quotes in the response")
		return
	}

	quote = qs[0]
	return
}

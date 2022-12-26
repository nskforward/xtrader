package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nskforward/xtrader/internal/quote"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetQuotes(ctx context.Context, cfg Config, symbol string) (quotes chan quote.Quote, err error) {
	url := fmt.Sprintf("%s/md/3.0/feed/%s", cfg.Addr, symbol)
	req, err2 := http.NewRequest("GET", url, nil)
	if err2 != nil {
		err = err2
		return
	}
	req.WithContext(ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", NewToken(cfg)}, " "))
	req.Header.Add("Accept", "application/x-json-stream")
	resp, err2 := http.DefaultClient.Do(req)
	if err2 != nil {
		err = err2
		return
	}

	if resp.StatusCode > 399 {
		data, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			err = err2
			return
		}
		_ = resp.Body.Close()
		err = fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
		return
	}

	quotes = make(chan quote.Quote, 1)
	dataChannel := make(chan json.RawMessage, 1)

	go func() {
		defer close(quotes)

	LOOP:
		for {
			select {
			case <-ctx.Done():
				return

			case data, ok := <-dataChannel:
				if !ok {
					return
				}
				var q Quote
				err := json.Unmarshal(data, &q)
				if err != nil {
					fmt.Println("[error] unmarshal quote:", err)
					fmt.Println(string(data))
					continue LOOP
				}
				if q.Event != "" {
					fmt.Println("[info] quote event:", q.Event)
					continue LOOP
				}
				quotes <- q
			}
		}
	}()

	go func() {
		defer resp.Body.Close()
		defer close(dataChannel)

		decoder := json.NewDecoder(resp.Body)

		for {
			var jsonString json.RawMessage
			err := decoder.Decode(&jsonString)
			if err == io.EOF {
				fmt.Println("[info] stream stopped:", symbol)
				return
			}
			if err != nil {
				fmt.Println("[error] cannot decode quote response body:", err)
				return
			}
			if len(jsonString) == 0 {
				continue
			}

			select {
			case dataChannel <- jsonString:
			default: // skip
			}
		}
	}()

	return
}

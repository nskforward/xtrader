package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nskforward/xtrader/internal/broker"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func WatchQuotes(ctx context.Context, cfg Config, symbol string) (quotes chan broker.Quote, err2 error) {
	url := fmt.Sprintf("%s/md/3.0/feed/%s", cfg.Addr, symbol)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err2 = err
		return
	}
	req.WithContext(ctx)
	req.Header.Add("Authorization", strings.Join([]string{"Bearer", NewToken(cfg)}, " "))
	req.Header.Add("Accept", "application/x-json-stream")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		err2 = err
		return
	}

	if resp.StatusCode > 399 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			err2 = err
			return
		}
		resp.Body.Close()
		err2 = fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
		return
	}

	quotes = make(chan broker.Quote, 1)
	dataChannel := make(chan json.RawMessage, 1)

	go func() {
		defer close(quotes)

		var tmp Quote
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

	LOOP:
		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				if tmp.Timestamp > 0 {
					select {
					case quotes <- tmp:
						tmp.Timestamp = 0
					default:
						<-quotes
					}
				}

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
				tmp = q
			}
		}
	}()

	go func() {
		defer resp.Body.Close()
		defer close(dataChannel)

		//r := bufio.NewReader(resp.Body)
		decoder := json.NewDecoder(resp.Body)

		for {
			//data, err := r.ReadBytes('}')
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

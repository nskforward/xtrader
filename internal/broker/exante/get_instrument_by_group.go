package exante

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetInstrumentByGroup(ctx context.Context, cfg Config, group string, f func(inst Instrument) bool) error {
	url := fmt.Sprintf("%s/md/3.0/groups/%s", cfg.Addr, group)
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

	if resp.StatusCode > 399 {
		data, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("bad http response code: %s: %s", resp.Status, string(data))
	}

	decoder := json.NewDecoder(resp.Body)

	// read open bracket
	_, err = decoder.Token()
	if err != nil {
		return err
	}

	var ins Instrument

LOOP:
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			if !decoder.More() {
				break LOOP
			}
			err := decoder.Decode(&ins)
			if err != nil {
				return err
			}
			if !f(ins) {
				return nil
			}
		}
	}

	// read closing bracket
	_, err = decoder.Token()
	if err != nil {
		return err
	}
	return nil
}

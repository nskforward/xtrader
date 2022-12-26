package main

import (
	"context"
	"github.com/nskforward/xtrader/internal/broker/exante"
	"github.com/nskforward/xtrader/internal/config"
	"github.com/nskforward/xtrader/internal/httpclient"
	"github.com/nskforward/xtrader/pkg/logger"
	"os"
	"time"
)

/*
get instrument INTC.NASDAQ
{"optionData":null,"name":"Intel Corporation","symbolId":"INTC.NASDAQ","description":"Intel Corporation - Common Stock","icon":"https://circles-all.s3.eu-central-1.amazonaws.com/STOCK/bkN0HaNeKvdxRBGv1n2teS9MsYUmsSG.png","underlyingSymbolId":null,"country":"US","identifiers":{"ISIN":"US4581401001","FIGI":"BBG000C0G1D1","RIC":"INTC.OQ","SEDOL":"2463247"},"exchange":"NASDAQ","symbolType":"STOCK","currency":"USD","minPriceIncrement":"0.01","ticker":"INTC","expiration":null,"group":null}

get instrument INTC.CBOE.30Z2022.C28
{"optionData":{"optionGroupId":"INTC.CBOE.30Z2022.C*","strikePrice":"28","optionRight":"CALL"},"name":"Intel","symbolId":"INTC.CBOE.30Z2022.C28","description":"Option on Intel 30 Dec 2022 CALL 28","icon":null,"underlyingSymbolId":"INTC.NASDAQ","country":"US","identifiers":null,"exchange":"CBOE","symbolType":"OPTION","currency":"USD","minPriceIncrement":"0.01","ticker":"INTC","expiration":1672434000000,"group":"INTC"}


*/

func main() {
	l := logger.New(true, true, os.Stdout, os.Stderr)
	httpclient.InitDefault()

	l.Info("service started")
	defer l.Info("service stopped")

	cfg := config.Get()
	l.Info("broker:", cfg.Broker.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	/*
		order := exante.Order{
			AccountID:  "PIX0219.007",
			ClientTag:  "robot v0.0.1-alpha.1",
			SymbolID:   "INTC.CBOE.30Z2022.C28",
			Side:       "sell",
			Quantity:   "1",
			OrderType:  "limit",
			LimitPrice: "0.5",
			Duration:   "good_till_cancel",
		}

		respArr, err := exante.PlaceOrder(ctx, cfg.Broker.Exante, order)
		if err != nil {
			l.Error(err)
		}

		for _, resp := range respArr {
			l.Info("status:", resp.OrderState.Status)
			l.Info("reason:", resp.OrderState.Reason)
		}
	*/

	err := exante.GetActiveOrders(ctx, cfg.Broker.Exante, "INTC.CBOE.30Z2022.C28")
	if err != nil {
		l.Error(err)
	}

	// 63a1c7fb-d0bb-47d1-8b6b-a0b32d8c4213
}

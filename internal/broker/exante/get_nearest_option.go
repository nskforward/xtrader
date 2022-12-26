package exante

import (
	"context"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func GetNearestOption(ctx context.Context, cfg Config, underlyingSymbolId, optionRight string, start time.Time, minDistance float64) (instrument Instrument, err error) {
	items := strings.Split(underlyingSymbolId, ".")
	if len(items) != 2 {
		err = fmt.Errorf("cannot extract group from underlyingSymbolId")
		return
	}
	groupID := items[0]

	// get last price
	last, err2 := GetLastPrice(ctx, cfg, underlyingSymbolId)
	if err2 != nil {
		err = err2
		return
	}
	priceBid, _, err2 := last.Bid_()
	if err2 != nil {
		err = err2
		return
	}
	priceAsk, _, err2 := last.Ask_()
	if err2 != nil {
		err = err2
		return
	}
	mid := (priceBid + priceAsk) / 2
	if optionRight == "CALL" {
		mid = mid + minDistance
	} else {
		mid = mid - minDistance
	}

	var tmp Instrument
	var tmpDelta float64

	startUnix := start.UnixMilli()

	err2 = GetInstrumentByGroup(ctx, cfg, groupID, func(in Instrument) bool {
		if in.SymbolType != "OPTION" {
			return true
		}

		if in.OptionData.OptionRight != optionRight {
			return true
		}

		if in.Expiration < startUnix {
			return true
		}

		strike, err := strconv.ParseFloat(in.OptionData.StrikePrice, 64)
		if err != nil {
			fmt.Println("[error] cannot parse strike price:", err)
			return false
		}

		if tmp.Expiration == 0 {
			tmp = in
			tmpDelta = math.Abs(mid - strike)
			return true
		}

		if in.Expiration < tmp.Expiration {
			tmp = in
			tmpDelta = math.Abs(mid - strike)
			return true
		}

		delta := math.Abs(mid - strike)
		if delta < tmpDelta {
			tmp = in
			tmpDelta = delta
			return true
		}

		return true
	})

	if err2 != nil {
		err = err2
		return
	}

	if tmp.Expiration == 0 {
		err = fmt.Errorf("no options found")
		return
	}

	instrument = tmp
	return
}

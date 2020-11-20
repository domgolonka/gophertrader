package trading

import (
	"context"
	"gophertrader/gctrpc"
)

func (t Gctrpc) CandleHistory(ctx context.Context, exchange, assetType, start, end string, exrequest, sync, usedb, filltrades, force bool, pair *gctrpc.CurrencyPair) error {
	resp, err := t.GetHistoricCandles(ctx, &gctrpc.GetHistoricCandlesRequest{
		Exchange:              exchange,
		Pair:                  pair,
		AssetType:             assetType,
		Start:                 start,
		End:                   end,
		TimeInterval:          0,
		ExRequest:             exrequest,
		Sync:                  sync,
		UseDb:                 usedb,
		FillMissingWithTrades: filltrades,
		Force:                 force,
	})
	if err != nil {
		return err
	}
	resp.GetCandle()

	return nil
}

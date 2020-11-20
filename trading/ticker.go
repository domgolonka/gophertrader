package trading

import (
	"context"
	"gophertrader/gctrpc"
	"io"
	"log"
)

func (t Gctrpc) TickerStream(ctx context.Context, exchange, assetType string, pair *gctrpc.CurrencyPair) error {
	stream, err := t.GetTickerStream(ctx, &gctrpc.GetTickerStreamRequest{
		Exchange: exchange,
		Pair: pair,
		AssetType: assetType,
	})
	if err != nil {
		return err
	}

	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error : %v",err)
		}
		feature.
	}
	return nil
}

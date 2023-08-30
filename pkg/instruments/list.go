package instruments

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"robinhood/internal"
	"robinhood/pkg/models/client/robinhood"
)

func ListInstruments(ctx context.Context, host string, bearerToken string) (*robinhood.Instruments, error) {
	var results *robinhood.Instruments
	body, err := internal.NewApiRequestGetData(ctx, host, "instruments/", bearerToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &results)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarssahalling:: %v", err))
	}

	return results, nil

}

func ListInstrumentsByListId(ctx context.Context, host string, bearerToken string) (*robinhood.ListResultsByType, error) {
	var results *robinhood.ListResultsByType
	body, err := internal.NewApiRequestGetData(ctx, host, "midlands/lists/items/?list_id=eddbebe5-34cc-4df1-953c-d3e3cb55bc19&local_midnight=2023-08-29T07%3A00%3A00.000Z", bearerToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &results)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarssahalling:: %v", err))
	}

	return results, nil

}

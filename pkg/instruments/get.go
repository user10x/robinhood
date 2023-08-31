package instruments

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"robinhood/internal"
	"robinhood/pkg/models/client/robinhood"
)

func GetRatingById(ctx context.Context, host, bearerToken, id string) (*robinhood.ListResultsByType, error) {
	var results *robinhood.ListResultsByType
	body, err := internal.NewApiRequestGetData(ctx, host, fmt.Sprintf("midlands/ratings/%s/", id), bearerToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &results)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarssahalling:: %v", err))
	}

	return results, nil

}

func GetInstrumentsBySymbol(ctx context.Context, host, bearerToken, symbol string) (*robinhood.ListResultsByType, error) {
	var results *robinhood.ListResultsByType
	body, err := internal.NewApiRequestGetData(ctx, host, fmt.Sprintf("midlands/ratings/4966fbab-38d6-4eaa-969f-e26e262d8e20/", symbol), bearerToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &results)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarssahalling:: %v", err))
	}

	return results, nil

}

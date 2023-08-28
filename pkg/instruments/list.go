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

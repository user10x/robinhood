package feed

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"robinhood/internal"
	"robinhood/pkg/models/client/robinhood"
)

func ListFeeds(ctx context.Context, host string, bearerToken string) ([]robinhood.Feed, error) {
	var feeds *robinhood.Feeds
	body, err := internal.NewApiRequestGetData(ctx, host, "feed/", bearerToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &feeds)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarsahaling: %v", err))
	}

	return feeds.Results, nil

}

package orders

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"robinhood/internal"
	"robinhood/pkg/models/client/robinhood"
)

func ListOrders(ctx context.Context, host string, suffix, bearerToken string) ([]robinhood.Order, error) {
	var orders *robinhood.Orders
	body, err := internal.NewApiRequestGetData(ctx, host, suffix, bearerToken)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &orders)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarsahaling: %v", err))
	}

	return orders.Results, nil

}

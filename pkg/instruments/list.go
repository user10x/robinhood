package instruments

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"robinhood/pkg/models/client/robinhood"
	"time"
)

func setHeaders(req *http.Request, bearerToken string) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	req.Header.Set("agent", "Robinhood/823 (iPhone; iOS 7.1.2; Scale/2.00)")
	req.Header.Set("Content-Type", "application/json")
	return req
}

func ListInstruments(ctx context.Context, host string, bearerToken string) (*robinhood.Instruments, error) {

	client := &http.Client{
		Timeout: time.Duration(100 * time.Second),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%s/%s", host, "instruments/"), nil)
	if err != nil {
		return nil, err
	}
	req = setHeaders(req, bearerToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 300 {
		return nil, errors.New(fmt.Sprintf("error status code %d", resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error while reading the response bytes: %v", err))
	}
	var results *robinhood.Instruments
	err = json.Unmarshal(body, &results)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error unmarssahalling:: %v", err))
	}

	return results, nil

}

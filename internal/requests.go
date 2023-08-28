package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func setHeaders(req *http.Request, bearerToken string) *http.Request {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	req.Header.Set("agent", "Robinhood/823 (iPhone; iOS 7.1.2; Scale/2.00)")
	req.Header.Set("Content-Type", "application/json")
	return req
}

func NewApiRequestGetData(ctx context.Context, host string, suffix string, bearerToken string) ([]byte, error) {
	client := &http.Client{}

	path := fmt.Sprintf("%s/%s", host, suffix)
	req, err := http.NewRequestWithContext(ctx, "GET", path, nil)
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

	return body, nil

}

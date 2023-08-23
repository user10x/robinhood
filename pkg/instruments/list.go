package instruments

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"robinhood/pkg/models/client/robinhood"
	"robinhood/pkg/storage/firestore"
	"time"
)

func ListInstruments(ctx context.Context, host string, bearerToken string) ([]string, error) {

	inst := make([]string, 0)
	client := &http.Client{
		Timeout: time.Duration(100 * time.Second),
	}

	req, err := http.NewRequestWithContext(ctx, "GET", "https://api.robinhood.com/instruments/", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))

	req.Header.Set("agent", "Robinhood/823 (iPhone; iOS 7.1.2; Scale/2.00)")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	var results *robinhood.Instruments
	err = json.Unmarshal(body, &results)

	if err != nil {
		os.Exit(1)
	}
	
	storage := firestore.NewFireStore("States", "bc-roc-poc")
	storage.Insert(ctx, results)

	return inst, nil

}

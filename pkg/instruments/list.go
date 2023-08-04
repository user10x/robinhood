package instruments

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Results struct {
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	Id                                 string      `json:"id"`
	Url                                string      `json:"url"`
	Quote                              string      `json:"quote"`
	Fundamentals                       string      `json:"fundamentals"`
	Splits                             string      `json:"splits"`
	State                              string      `json:"state"`
	Market                             string      `json:"market"`
	SimpleName                         string      `json:"simple_name"`
	Name                               string      `json:"name"`
	Tradeable                          bool        `json:"tradeable"`
	Tradability                        string      `json:"tradability"`
	Symbol                             string      `json:"symbol"`
	BloombergUnique                    string      `json:"bloomberg_unique"`
	MarginInitialRatio                 string      `json:"margin_initial_ratio"`
	MaintenanceRatio                   string      `json:"maintenance_ratio"`
	Country                            string      `json:"country"`
	DayTradeRatio                      string      `json:"day_trade_ratio"`
	ListDate                           string      `json:"list_date"`
	MinTickSize                        interface{} `json:"min_tick_size"`
	Type                               string      `json:"type"`
	TradableChainId                    interface{} `json:"tradable_chain_id"`
	RhsTradability                     string      `json:"rhs_tradability"`
	FractionalTradability              string      `json:"fractional_tradability"`
	DefaultCollarFraction              string      `json:"default_collar_fraction"`
	IpoAccessStatus                    interface{} `json:"ipo_access_status"`
	IpoAccessCobDeadline               interface{} `json:"ipo_access_cob_deadline"`
	IpoS1Url                           interface{} `json:"ipo_s1_url"`
	IpoRoadshowUrl                     interface{} `json:"ipo_roadshow_url"`
	IsSpac                             bool        `json:"is_spac"`
	IsTest                             bool        `json:"is_test"`
	IpoAccessSupportsDsp               bool        `json:"ipo_access_supports_dsp"`
	ExtendedHoursFractionalTradability bool        `json:"extended_hours_fractional_tradability"`
	InternalHaltReason                 string      `json:"internal_halt_reason"`
	InternalHaltDetails                string      `json:"internal_halt_details"`
	InternalHaltSessions               interface{} `json:"internal_halt_sessions"`
	InternalHaltStartTime              interface{} `json:"internal_halt_start_time"`
	InternalHaltEndTime                interface{} `json:"internal_halt_end_time"`
	InternalHaltSource                 string      `json:"internal_halt_source"`
	AllDayTradability                  string      `json:"all_day_tradability"`
}

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
	var results *Results
	err = json.Unmarshal(body, &results)

	if err != nil {

	}

	for _, v := range results.Results {
		fmt.Println(v)
	}

	return inst, nil

}

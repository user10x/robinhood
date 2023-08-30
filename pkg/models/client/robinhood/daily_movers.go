package robinhood

import "time"

type DailyMovers struct {
	Results []struct {
		CreatedAt                  time.Time   `json:"created_at"`
		Id                         string      `json:"id"`
		ListId                     string      `json:"list_id"`
		ObjectId                   string      `json:"object_id"`
		ObjectType                 string      `json:"object_type"`
		OwnerType                  string      `json:"owner_type"`
		UpdatedAt                  time.Time   `json:"updated_at"`
		Weight                     string      `json:"weight"`
		OpenPrice                  interface{} `json:"open_price"`
		OpenPriceDirection         interface{} `json:"open_price_direction"`
		MarketCap                  float64     `json:"market_cap"`
		Name                       string      `json:"name"`
		OpenPositions              int         `json:"open_positions"`
		Symbol                     string      `json:"symbol"`
		UkTradability              string      `json:"uk_tradability"`
		UsTradability              string      `json:"us_tradability"`
		State                      string      `json:"state"`
		IpoAccessStatus            interface{} `json:"ipo_access_status"`
		TotalReturnPercentage      interface{} `json:"total_return_percentage"`
		Holdings                   bool        `json:"holdings"`
		OneDayDollarChange         float64     `json:"one_day_dollar_change"`
		OneDayPercentChange        float64     `json:"one_day_percent_change"`
		OneDayRollingDollarChange  float64     `json:"one_day_rolling_dollar_change"`
		OneDayRollingPercentChange float64     `json:"one_day_rolling_percent_change"`
		Price                      float64     `json:"price"`
	} `json:"results"`
}

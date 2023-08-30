package robinhood

import "time"

type Instruments struct {
	Next     string       `json:"next"`
	Previous interface{}  `json:"previous"`
	Results  []Instrument `json:"results"`
}

type Instrument struct {
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

type ListResultsByType struct {
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

package robinhood

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

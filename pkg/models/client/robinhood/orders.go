package robinhood

import "time"

type Orders struct {
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Order     `json:"results"`
}

type Order struct {
	Id                     string      `json:"id"`
	RefId                  string      `json:"ref_id"`
	Url                    string      `json:"url"`
	Account                string      `json:"account"`
	UserUuid               string      `json:"user_uuid"`
	Position               string      `json:"position"`
	Cancel                 interface{} `json:"cancel"`
	Instrument             string      `json:"instrument"`
	InstrumentId           string      `json:"instrument_id"`
	CumulativeQuantity     string      `json:"cumulative_quantity"`
	AveragePrice           *string     `json:"average_price"`
	Fees                   string      `json:"fees"`
	State                  string      `json:"state"`
	PendingCancelOpenAgent interface{} `json:"pending_cancel_open_agent"`
	Type                   string      `json:"type"`
	Side                   string      `json:"side"`
	TimeInForce            string      `json:"time_in_force"`
	Trigger                string      `json:"trigger"`
	Price                  *string     `json:"price"`
	StopPrice              interface{} `json:"stop_price"`
	Quantity               string      `json:"quantity"`
	RejectReason           interface{} `json:"reject_reason"`
	CreatedAt              time.Time   `json:"created_at"`
	UpdatedAt              time.Time   `json:"updated_at"`
	LastTransactionAt      time.Time   `json:"last_transaction_at"`
	Executions             []struct {
		Price                  string      `json:"price"`
		Quantity               string      `json:"quantity"`
		RoundedNotional        string      `json:"rounded_notional"`
		SettlementDate         string      `json:"settlement_date"`
		Timestamp              time.Time   `json:"timestamp"`
		Id                     string      `json:"id"`
		IpoAccessExecutionRank interface{} `json:"ipo_access_execution_rank"`
		TradeExecutionDate     *string     `json:"trade_execution_date"`
	} `json:"executions"`
	ExtendedHours           bool        `json:"extended_hours"`
	MarketHours             string      `json:"market_hours"`
	OverrideDtbpChecks      bool        `json:"override_dtbp_checks"`
	OverrideDayTradeChecks  bool        `json:"override_day_trade_checks"`
	ResponseCategory        *string     `json:"response_category"`
	StopTriggeredAt         interface{} `json:"stop_triggered_at"`
	LastTrailPrice          interface{} `json:"last_trail_price"`
	LastTrailPriceUpdatedAt interface{} `json:"last_trail_price_updated_at"`
	LastTrailPriceSource    interface{} `json:"last_trail_price_source"`
	DollarBasedAmount       *struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
		CurrencyId   string `json:"currency_id"`
	} `json:"dollar_based_amount"`
	TotalNotional *struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
		CurrencyId   string `json:"currency_id"`
	} `json:"total_notional"`
	ExecutedNotional *struct {
		Amount       string `json:"amount"`
		CurrencyCode string `json:"currency_code"`
		CurrencyId   string `json:"currency_id"`
	} `json:"executed_notional"`
	InvestmentScheduleId         interface{} `json:"investment_schedule_id"`
	IsIpoAccessOrder             bool        `json:"is_ipo_access_order"`
	IpoAccessCancellationReason  *string     `json:"ipo_access_cancellation_reason"`
	IpoAccessLowerCollaredPrice  interface{} `json:"ipo_access_lower_collared_price"`
	IpoAccessUpperCollaredPrice  interface{} `json:"ipo_access_upper_collared_price"`
	IpoAccessUpperPrice          interface{} `json:"ipo_access_upper_price"`
	IpoAccessLowerPrice          interface{} `json:"ipo_access_lower_price"`
	IsIpoAccessPriceFinalized    bool        `json:"is_ipo_access_price_finalized"`
	IsVisibleToUser              bool        `json:"is_visible_to_user"`
	HasIpoAccessCustomPriceLimit bool        `json:"has_ipo_access_custom_price_limit"`
	IsPrimaryAccount             bool        `json:"is_primary_account"`
	OrderFormVersion             int         `json:"order_form_version"`
	PresetPercentLimit           interface{} `json:"preset_percent_limit"`
	OrderFormType                *string     `json:"order_form_type"`
}

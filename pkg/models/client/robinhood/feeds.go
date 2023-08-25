package robinhood

type Feeds struct {
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Feed      `json:"results"`
}

type Feed struct {
	DisplayLabel string   `json:"display_label"`
	Category     string   `json:"category"`
	Templates    []string `json:"templates"`
	Contents     []struct {
		ContentType string `json:"content_type"`
		Data        struct {
			Label         string `json:"label"`
			IndicatorType string `json:"indicator_type"`
			IndicatorKey  string `json:"indicator_key"`
		} `json:"data"`
		ID               string      `json:"id"`
		Reason           interface{} `json:"reason"`
		InstrumentID     interface{} `json:"instrument_id"`
		InstrumentSector interface{} `json:"instrument_sector"`
	} `json:"contents"`
	URL                    string      `json:"url"`
	Description            interface{} `json:"description"`
	RankingVersion         string      `json:"ranking_version"`
	ID                     string      `json:"id"`
	LogoAssetName          interface{} `json:"logo_asset_name"`
	DisplayLabelInfoAction interface{} `json:"display_label_info_action"`
	FeedType               interface{} `json:"feed_type"`
	FeedLocation           interface{} `json:"feed_location"`
}

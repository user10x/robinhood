package robinhood

import "time"

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
			Label              string    `json:"label,omitempty"`
			IndicatorType      string    `json:"indicator_type,omitempty"`
			IndicatorKey       string    `json:"indicator_key,omitempty"`
			Id                 string    `json:"id,omitempty"`
			DisplayName        string    `json:"display_name,omitempty"`
			OwnerType          string    `json:"owner_type,omitempty"`
			ItemCount          int       `json:"item_count,omitempty"`
			IsBadged           bool      `json:"is_badged,omitempty"`
			Source             string    `json:"source,omitempty"`
			Title              string    `json:"title,omitempty"`
			PublishedAt        time.Time `json:"published_at,omitempty"`
			RelatedInstruments []struct {
				InstrumentId string      `json:"instrument_id"`
				Symbol       string      `json:"symbol"`
				Name         string      `json:"name"`
				Sector       interface{} `json:"sector"`
				SimpleName   string      `json:"simple_name"`
			} `json:"related_instruments,omitempty"`
			Url      *string `json:"url,omitempty"`
			Feedback struct {
				PositiveCount int `json:"positive_count"`
			} `json:"feedback,omitempty"`
			Media struct {
				Url      string `json:"url"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Mimetype string `json:"mimetype"`
			} `json:"media,omitempty"`
			PreviewMedia interface{} `json:"preview_media"`
			PreviewText  string      `json:"preview_text,omitempty"`
			IsEmbedded   bool        `json:"is_embedded,omitempty"`
			LogoHexCode  *string     `json:"logo_hex_code,omitempty"`
			Authors      interface{} `json:"authors,omitempty"`
			Popularity   int         `json:"popularity,omitempty"`
			Thumbnail    struct {
				Url      string `json:"url"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Mimetype string `json:"mimetype"`
			} `json:"thumbnail,omitempty"`
			Captions     string `json:"captions,omitempty"`
			InstrumentId string `json:"instrument_id,omitempty"`
			Symbol       string `json:"symbol,omitempty"`
			Name         string `json:"name,omitempty"`
			Sector       string `json:"sector,omitempty"`
			SimpleName   string `json:"simple_name,omitempty"`
		} `json:"data"`
		Id               string      `json:"id"`
		Reason           *string     `json:"reason"`
		InstrumentId     *string     `json:"instrument_id"`
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

package robinhood

import "time"

type Popular struct {
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		DisplayLabel string   `json:"display_label"`
		Category     string   `json:"category"`
		Templates    []string `json:"templates"`
		Contents     []struct {
			ContentType string `json:"content_type"`
			Data        struct {
				Label         string `json:"label,omitempty"`
				IndicatorType string `json:"indicator_type,omitempty"`
				IndicatorKey  string `json:"indicator_key,omitempty"`
				Id            string `json:"id,omitempty"`
				DisplayName   string `json:"display_name,omitempty"`
				ImageUrls     struct {
					Circle281    string `json:"circle_28:1"`
					Circle2815   string `json:"circle_28:1.5"`
					Circle282    string `json:"circle_28:2"`
					Circle283    string `json:"circle_28:3"`
					Circle641    string `json:"circle_64:1"`
					Circle6415   string `json:"circle_64:1.5"`
					Circle642    string `json:"circle_64:2"`
					Circle643    string `json:"circle_64:3"`
					HeaderMob1   string `json:"header_mob:1"`
					HeaderMob15  string `json:"header_mob:1.5"`
					HeaderMob2   string `json:"header_mob:2"`
					HeaderMob3   string `json:"header_mob:3"`
					HeaderWeb1   string `json:"header_web:1"`
					HeaderWeb15  string `json:"header_web:1.5"`
					HeaderWeb2   string `json:"header_web:2"`
					HeaderWeb3   string `json:"header_web:3"`
					Portrait481  string `json:"portrait_48:1"`
					Portrait4815 string `json:"portrait_48:1.5"`
					Portrait482  string `json:"portrait_48:2"`
					Portrait483  string `json:"portrait_48:3"`
				} `json:"image_urls,omitempty"`
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
		Url                    string      `json:"url"`
		Description            *string     `json:"description"`
		RankingVersion         string      `json:"ranking_version"`
		Id                     string      `json:"id"`
		LogoAssetName          *string     `json:"logo_asset_name"`
		DisplayLabelInfoAction interface{} `json:"display_label_info_action"`
		FeedType               interface{} `json:"feed_type"`
		FeedLocation           interface{} `json:"feed_location"`
	} `json:"results"`
}

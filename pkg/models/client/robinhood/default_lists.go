package robinhood

import "time"

type T struct {
	Id       string `json:"id"`
	Migrated bool   `json:"migrated"`
	Results  []struct {
		ChildSortDirection string    `json:"child_sort_direction"`
		ChildSortOrder     string    `json:"child_sort_order"`
		CreatedAt          time.Time `json:"created_at"`
		DisplayDescription *string   `json:"display_description"`
		DisplayName        string    `json:"display_name"`
		Id                 string    `json:"id"`
		OwnerType          string    `json:"owner_type"`
		ParentLists        []struct {
			BackgroundColor string `json:"background_color"`
			ChildInfo       struct {
				ChildType string        `json:"child_type"`
				Children  []interface{} `json:"children"`
			} `json:"child_info"`
			CreatedAt                  time.Time   `json:"created_at"`
			DefaultExpanded            bool        `json:"default_expanded"`
			DisplayDescription         string      `json:"display_description"`
			MarkdownDisplayDescription interface{} `json:"markdown_display_description"`
			DisplayName                string      `json:"display_name"`
			ForegroundColor            string      `json:"foreground_color"`
			Icon                       string      `json:"icon"`
			ItemCount                  int         `json:"item_count"`
			Id                         string      `json:"id"`
			OwnerType                  string      `json:"owner_type"`
			ReadPermission             string      `json:"read_permission"`
			UpdatedAt                  time.Time   `json:"updated_at"`
			BackgroundColorToken       interface{} `json:"background_color_token"`
			BackgroundGradient         interface{} `json:"background_gradient"`
			Disclosure                 struct {
				Attributes []interface{} `json:"attributes"`
				Text       string        `json:"text"`
			} `json:"disclosure"`
			CollapsedDisclosure struct {
				Attributes []interface{} `json:"attributes"`
				Text       string        `json:"text"`
			} `json:"collapsed_disclosure"`
			Truncated bool `json:"truncated"`
		} `json:"parent_lists"`
		ReadPermission     string    `json:"read_permission"`
		UpdatedAt          time.Time `json:"updated_at"`
		AllowedObjectTypes []string  `json:"allowed_object_types"`
		IconEmoji          string    `json:"icon_emoji,omitempty"`
		Owner              string    `json:"owner,omitempty"`
		ItemCount          int       `json:"item_count"`
		ChildInfo          struct {
			ChildType string        `json:"child_type"`
			Children  []interface{} `json:"children"`
		} `json:"child_info"`
		Followed                     bool          `json:"followed"`
		DefaultExpanded              bool          `json:"default_expanded"`
		RelatedLists                 []interface{} `json:"related_lists"`
		BackgroundColor              string        `json:"background_color,omitempty"`
		BackgroundColorToken         interface{}   `json:"background_color_token"`
		BackgroundGradient           interface{}   `json:"background_gradient"`
		ForegroundColor              string        `json:"foreground_color,omitempty"`
		ImmediatelyVisibleDisclosure bool          `json:"immediately_visible_disclosure,omitempty"`
		MarkdownDisplayDescription   *string       `json:"markdown_display_description,omitempty"`
		Icon                         *string       `json:"icon,omitempty"`
		Disclosure                   struct {
			Attributes []struct {
				Location int    `json:"location"`
				Length   int    `json:"length"`
				Link     string `json:"link"`
				Style    string `json:"style"`
			} `json:"attributes"`
			Text string `json:"text"`
		} `json:"disclosure,omitempty"`
		CollapsedDisclosure struct {
			Attributes []interface{} `json:"attributes"`
			Text       string        `json:"text"`
		} `json:"collapsed_disclosure,omitempty"`
		Truncated bool `json:"truncated,omitempty"`
	} `json:"results"`
}

package webhook

// Message represents a Discord webhook message
type Message struct {
	Content   string   `json:"content,omitempty"`
	Username  string   `json:"username,omitempty"`
	AvatarURL string   `json:"avatar_url,omitempty"`
	TTS       bool     `json:"tts,omitempty"`
	Embeds    []*Embed `json:"embeds,omitempty"`
}

// Embed represents a Discord embed
type Embed struct {
	Title       string         `json:"title,omitempty"`
	Description string         `json:"description,omitempty"`
	URL         string         `json:"url,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
	Color       int            `json:"color,omitempty"`
	Footer      *EmbedFooter   `json:"footer,omitempty"`
	Image       *EmbedImage    `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Author      *EmbedAuthor   `json:"author,omitempty"`
	Fields      []*EmbedField  `json:"fields,omitempty"`
}

// EmbedFooter represents a Discord embed footer
type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

// EmbedImage represents a Discord embed image
type EmbedImage struct {
	URL string `json:"url"`
}

// EmbedThumbnail represents a Discord embed thumbnail
type EmbedThumbnail struct {
	URL string `json:"url"`
}

// EmbedAuthor represents a Discord embed author
type EmbedAuthor struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// EmbedField represents a Discord embed field
type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

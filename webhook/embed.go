package webhook

import (
	"time"
)

// NewEmbed creates a new Discord embed
func NewEmbed() *Embed {
	return &Embed{
		Fields: []*EmbedField{},
	}
}

// SetTitle sets the embed title
func (e *Embed) SetTitle(title string) *Embed {
	result := *e
	result.Title = title
	return &result
}

// SetDescription sets the embed description
func (e *Embed) SetDescription(description string) *Embed {
	result := *e
	result.Description = description
	return &result
}

// SetURL sets the embed URL
func (e *Embed) SetURL(url string) *Embed {
	result := *e
	result.URL = url
	return &result
}

// SetTimestamp sets the embed timestamp
// If t is nil, the current time is used
func (e *Embed) SetTimestamp(t *time.Time) *Embed {
	result := *e

	if t == nil {
		now := time.Now().UTC()
		t = &now
	}

	result.Timestamp = t.Format(time.RFC3339)
	return &result
}

// SetColor sets the embed color
func (e *Embed) SetColor(color int) *Embed {
	result := *e
	result.Color = color
	return &result
}

// AddField adds a field to the embed
func (e *Embed) AddField(name, value string, inline bool) *Embed {
	result := *e

	result.Fields = make([]*EmbedField, len(e.Fields))
	copy(result.Fields, e.Fields)

	result.Fields = append(result.Fields, &EmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	})

	return &result
}

// SetFooter sets the embed footer
func (e *Embed) SetFooter(text string, iconURL string) *Embed {
	result := *e
	result.Footer = &EmbedFooter{
		Text:    text,
		IconURL: iconURL,
	}
	return &result
}

// SetImage sets the embed image
func (e *Embed) SetImage(url string) *Embed {
	result := *e
	result.Image = &EmbedImage{
		URL: url,
	}
	return &result
}

// SetThumbnail sets the embed thumbnail
func (e *Embed) SetThumbnail(url string) *Embed {
	result := *e
	result.Thumbnail = &EmbedThumbnail{
		URL: url,
	}
	return &result
}

// SetAuthor sets the embed author
func (e *Embed) SetAuthor(name, url, iconURL string) *Embed {
	result := *e
	result.Author = &EmbedAuthor{
		Name:    name,
		URL:     url,
		IconURL: iconURL,
	}
	return &result
}

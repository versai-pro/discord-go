package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client is the main entry point for sending Discord webhooks
type Client struct {
	webhookURL  string
	httpClient  *http.Client
	rateLimiter *rateLimiter
}

// NewClient creates a new Discord webhook client
func NewClient(webhookURL string) *Client {
	return &Client{
		webhookURL:  webhookURL,
		httpClient:  &http.Client{Timeout: 10 * time.Second},
		rateLimiter: newRateLimiter(),
	}
}

// WithHTTPClient sets a custom HTTP client
func (c *Client) WithHTTPClient(httpClient *http.Client) *Client {
	c.httpClient = httpClient
	return c
}

// NewMessage creates a new Discord webhook message
func NewMessage() *Message {
	return &Message{
		Embeds: []*Embed{},
	}
}

// SetContent sets the message content
func (m *Message) SetContent(content string) *Message {
	result := *m
	result.Content = content
	return &result
}

// SetUsername sets the webhook username
func (m *Message) SetUsername(username string) *Message {
	result := *m
	result.Username = username
	return &result
}

// SetAvatarURL sets the webhook avatar URL
func (m *Message) SetAvatarURL(avatarURL string) *Message {
	result := *m
	result.AvatarURL = avatarURL
	return &result
}

// SetTTS enables or disables text-to-speech
func (m *Message) SetTTS(tts bool) *Message {
	result := *m
	result.TTS = tts
	return &result
}

// AddEmbed adds an embed to the message
func (m *Message) AddEmbed(embed *Embed) *Message {
	result := *m

	result.Embeds = make([]*Embed, len(m.Embeds))
	copy(result.Embeds, m.Embeds)

	result.Embeds = append(result.Embeds, embed)
	return &result
}

// Send sends a webhook message to Discord
func (c *Client) Send(message *Message) error {
	if err := c.rateLimiter.wait(); err != nil {
		return err
	}

	payload, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, c.webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "DiscordGoWebhook/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	c.rateLimiter.update(resp.Header)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return &HTTPError{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	return nil
}

// SendEmbed is a convenience method for sending a message with a single embed
func (c *Client) SendEmbed(embed *Embed) error {
	message := NewMessage().AddEmbed(embed)
	return c.Send(message)
}

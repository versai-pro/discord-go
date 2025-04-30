# Discord Webhook Go

[![Go Reference](https://pkg.go.dev/badge/github.com/versai-pro/discord-go.svg)](https://pkg.go.dev/github.com/versai-pro/discord-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/versai-pro/discord-go)](https://goreportcard.com/report/github.com/versai-pro/discord-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight, dependency-free Go library for sending Discord webhook messages with a fluent API for building embeds.

## Features

- **Simple API**: Easy-to-use methods for sending webhook messages
- **Fluent Embed Builder**: Method chaining for building rich embeds
- **Zero Dependencies**: Uses only Go's standard library
- **Rate Limit Handling**: Built-in support for Discord's rate limits
- **Immutable Design**: Thread-safe with copy-on-write pattern

## Installation

```bash
go get github.com/versai-pro/discord-go
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/versai-pro/discord-go/webhook"
)

func main() {
    // Create a new webhook client with your Discord webhook URL
    client := webhook.NewClient("https://discord.com/api/webhooks/your-webhook-id/your-webhook-token")

    // Create a simple embed
    embed := webhook.NewEmbed().
        SetTitle("Hello from Discord Go").
        SetDescription("This is a test message.").
        SetColor(0x00ff00) // Green color

    // Send the embed
    if err := client.SendEmbed(embed); err != nil {
        log.Fatalf("Failed to send webhook: %v", err)
    }

    fmt.Println("Webhook sent successfully!")
}
```

## Detailed Usage

### Creating a Webhook Client

```go
// Basic client
client := webhook.NewClient("https://discord.com/api/webhooks/your-webhook-id/your-webhook-token")

// Client with custom HTTP client
customHTTPClient := &http.Client{Timeout: 5 * time.Second}
client := webhook.NewClient("https://discord.com/api/webhooks/your-webhook-id/your-webhook-token").
    WithHTTPClient(customHTTPClient)
```

### Sending a Simple Message

```go
message := webhook.NewMessage().
    SetContent("This is a simple message").
    SetUsername("Custom Bot Name") // Optional: Override the webhook's default name

if err := client.Send(message); err != nil {
    log.Fatalf("Failed to send message: %v", err)
}
```

### Building and Sending an Embed

```go
// Create a new embed
embed := webhook.NewEmbed().
    SetTitle("Embed Title").
    SetDescription("This is an embed description.").
    SetURL("https://example.com"). // Optional: Add a URL to the title
    SetColor(0x3498db). // Blue color in hex
    SetTimestamp(nil) // Use current time

// Add fields to the embed
embed = embed.
    AddField("Regular Field", "This is a regular field", false).
    AddField("Inline Field 1", "This appears inline", true).
    AddField("Inline Field 2", "This also appears inline", true)

// Add footer, image, thumbnail, and author
embed = embed.
    SetFooter("Footer text", "https://example.com/footer-icon.png").
    SetImage("https://example.com/image.png").
    SetThumbnail("https://example.com/thumbnail.png").
    SetAuthor("Author Name", "https://example.com", "https://example.com/author-icon.png")

// Create a message with the embed
message := webhook.NewMessage().
    SetContent("Message with an embed").
    AddEmbed(embed)

// Send the message
if err := client.Send(message); err != nil {
    log.Fatalf("Failed to send message: %v", err)
}

// Alternatively, send just the embed
if err := client.SendEmbed(embed); err != nil {
    log.Fatalf("Failed to send embed: %v", err)
}
```

### Multiple Embeds

```go
// Create first embed
embed1 := webhook.NewEmbed().
    SetTitle("First Embed").
    SetDescription("This is the first embed.")

// Create second embed
embed2 := webhook.NewEmbed().
    SetTitle("Second Embed").
    SetDescription("This is the second embed.")

// Add both embeds to a message
message := webhook.NewMessage().
    AddEmbed(embed1).
    AddEmbed(embed2)

// Send the message with multiple embeds
if err := client.Send(message); err != nil {
    log.Fatalf("Failed to send message: %v", err)
}
```

### Error Handling

```go
// Send a message and handle potential errors
err := client.Send(message)
if err != nil {
    // Check for rate limit errors
    if rateLimitErr, ok := err.(*webhook.RateLimitError); ok {
        fmt.Printf("Rate limited! Retry after %v\n", rateLimitErr.RetryAfter)
        time.Sleep(rateLimitErr.RetryAfter)
        // Try again...
    }

    // Check for HTTP errors
    if httpErr, ok := err.(*webhook.HTTPError); ok {
        fmt.Printf("HTTP error %d: %s\n", httpErr.StatusCode, httpErr.Message)
        // Handle based on status code...
    }

    // Other errors
    fmt.Printf("Error: %v\n", err)
}
```

## Embed Color Reference

Discord uses integer color values. Here are some common colors you can use:

```go
// Discord Colors
const (
    ColorDefault = 0           // Default (invisible embed color)
    ColorBlue    = 0x3498db    // Blue
    ColorGreen   = 0x2ecc71    // Green
    ColorYellow  = 0xf1c40f    // Yellow
    ColorRed     = 0xe74c3c    // Red
    ColorPurple  = 0x9b59b6    // Purple
    ColorOrange  = 0xe67e22    // Orange
    ColorGrey    = 0x95a5a6    // Grey
    ColorNavy    = 0x34495e    // Navy
    ColorGold    = 0xf1c40f    // Gold
    ColorPink    = 0xe91e63    // Pink
)
```

## Discord Embed Limits

Discord enforces the following limits for embeds:

- Title: 256 characters
- Description: 4096 characters
- Fields: Up to 25 fields
- Field name: 256 characters
- Field value: 1024 characters
- Footer text: 2048 characters
- Author name: 256 characters
- Total embed characters: 6000 characters

## Rate Limiting

This library automatically handles Discord's rate limits. If you hit a rate limit, the library will return a `RateLimitError` with the `RetryAfter` duration. You can use this to wait before retrying.

## Example

Check out the [basic_webhook.go](examples/basic_webhook.go) example to see how to use this library.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## Acknowledgments

- [Discord API Documentation](https://discord.com/developers/docs/resources/webhook)
- [Discord Embed Visualizer](https://leovoel.github.io/embed-visualizer/) - Useful tool for previewing embeds

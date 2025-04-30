package main

import (
	"fmt"
	"log"
	"time"

	"github.com/versai-pro/discord-go/webhook"
)

func main() {
	webhookURL := "https://ptb.discord.com/api/webhooks/1367287466616422450/EZLSnIlUJpJcl8DFY1bPc-lvICtMtlMT8Akf_U51VIZX3sg97ECYHOos0cdkIcqgFg9z"

	client := webhook.NewClient(webhookURL)

	embed := webhook.NewEmbed().
		SetTitle("Hello from Discord Go Webhook").
		SetDescription("This is a test message sent using the Discord Go Webhook library.").
		SetColor(0x00ff00).
		SetTimestamp(nil).
		AddField("Field 1", "This is a regular field", false).
		AddField("Field 2", "This is an inline field", true).
		AddField("Field 3", "This is another inline field", true).
		SetFooter("Discord Go Webhook", "https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_Blue.png")

	message := webhook.NewMessage().
		SetContent("This is a test message").
		SetUsername("Go Webhook Bot").
		AddEmbed(embed)

	if err := client.Send(message); err != nil {
		log.Fatalf("Failed to send webhook: %v", err)
	}

	fmt.Println("Webhook sent successfully!")

	time.Sleep(1 * time.Second)

	simpleEmbed := webhook.NewEmbed().
		SetTitle("Simple Embed").
		SetDescription("This is a simple embed sent directly.")

	if err := client.SendEmbed(simpleEmbed); err != nil {
		log.Fatalf("Failed to send embed: %v", err)
	}

	fmt.Println("Embed sent successfully!")
}

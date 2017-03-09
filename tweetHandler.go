package main

import (
	"log"

	language "cloud.google.com/go/language/apiv1"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/net/context"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
)

type Handler struct {
	context context.Context
	client  language.Client
}

func NewHandler() *Handler {
	ctx := context.Background()
	client, _ := language.NewClient(ctx)
	return &Handler{context: ctx, client: *client}
}

func (h *Handler) Handle(tweet *twitter.Tweet) {
	sentimentScore := h.getSentiment(tweet.Text)

	sentimentText := "ERROR"
	if sentimentScore > 0 {
		sentimentText = "POSITIVE"
	} else if sentimentScore < 0 {
		sentimentText = "NEGATIVE"
	} else if sentimentScore == 0 {
		sentimentText = "NEUTRAL"
	}

	log.Printf("%s is %s with score %f", tweet.Text, sentimentText, sentimentScore)
}

func (h *Handler) getSentiment(text string) float32 {
	sentiment, err := h.client.AnalyzeSentiment(h.context, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
		EncodingType: languagepb.EncodingType_UTF8,
	})

	if err != nil {
		log.Println("Faied to get sentiment with error: ", err)
	}

	return sentiment.DocumentSentiment.Score
}

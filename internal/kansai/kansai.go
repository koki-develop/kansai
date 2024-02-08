package kansai

import (
	"context"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type Client struct {
	client *genai.Client
}

func New(ctx context.Context, apiKey string) (*Client, error) {
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	return &Client{client: client}, nil
}

func (c *Client) Close() error {
	return c.client.Close()
}

func (c *Client) GenerateContentStream(ctx context.Context, prompt string, callback func(p genai.Part) error) error {
	model := c.client.GenerativeModel("gemini-pro")
	iter := model.GenerateContentStream(ctx, genai.Text(prompt))

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				if err := callback(p); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

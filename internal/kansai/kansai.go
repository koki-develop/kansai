package kansai

import (
	"context"
	"fmt"

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

func (c *Client) Convert(ctx context.Context, s string, callback func(p genai.Part) error) error {
	prompt := fmt.Sprintf(`以下の $$$ 以降の文章を、できる限り元の文章の意味を保ったまま、自然な関西弁に変換してください。
なお、余計な解説や説明は一切不要です。変換後の文章のみを出力してください。
$$$
%s`, s)

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

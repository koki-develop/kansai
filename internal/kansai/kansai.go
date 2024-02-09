package kansai

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const (
	promptFormat = `以下の $$$ 以降のテキストを、元の意味を保ったまま自然な関西弁に変換してください。
テキストが日本語以外の言語の場合は、まずそれを日本語に翻訳した上で、関西弁に変換してください。
前置きや説明は一切不要です。変換後のテキストのみを提出してください。
$$$
%s`
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

func (c *Client) Convert(ctx context.Context, s string, callback func(p string) error) error {
	model := c.client.GenerativeModel("gemini-pro")

	iter := model.GenerateContentStream(ctx, genai.Text(fmt.Sprintf(promptFormat, s)))
	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		first := true
		for _, c := range resp.Candidates {
			for _, p := range c.Content.Parts {
				s := fmt.Sprint(p)
				if first {
					s = strings.TrimSpace(strings.TrimPrefix(s, "$$$")) // たまに先頭に $$$ が含まれることがあるので除去
					first = false
				}

				if err := callback(s); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

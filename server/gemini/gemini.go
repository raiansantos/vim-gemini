package gemini

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Client struct {
	gmClient *genai.Client
}

func (c *Client) extractAnswer(resp *genai.GenerateContentResponse) string {
	anwser := ""
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				anwser = fmt.Sprintf("%s%s\n\n", anwser, part)
			}
		}
	}
	return anwser
}

func (c *Client) DebugCode(ctx context.Context, filetype, code string) string {
	model := c.gmClient.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	res, err := cs.SendMessage(ctx, genai.Text(fmt.Sprintf("Could you debug this %s piece of code?\n\n%s", filetype, code)))
	if err != nil {
		log.Fatal(err)
	}
	return c.extractAnswer(res)
}

func (c *Client) ExplainCode(ctx context.Context, filetype, code string) string {
	model := c.gmClient.GenerativeModel("gemini-1.0-pro")
	cs := model.StartChat()

	res, err := cs.SendMessage(ctx, genai.Text(fmt.Sprintf("Could you explain me this %s piece of code?\n\n%s", filetype, code)))
	if err != nil {
		log.Fatal(err)
	}
	return c.extractAnswer(res)
}

func New(token string) (*Client, error) {
	ctx := context.Background()
	gmClient, err := genai.NewClient(ctx, option.WithAPIKey(token))
	if err != nil {
		return nil, err
	}

	return &Client{
		gmClient: gmClient,
	}, nil
}

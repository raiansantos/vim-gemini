package gemini

import (
	"context"
	"fmt"

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

func (c *Client) DebugCode(ctx context.Context, filetype, code string) (string, error) {
	model := c.gmClient.GenerativeModel("gemini-pro")
	cs := model.StartChat()

	message := fmt.Sprintf("Could you validate if this code is right, and if its not, debug this %s piece of code?\n\n%s", filetype, code)
	res, err := cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		return "", err
	}
	return c.extractAnswer(res), nil
}

func (c *Client) ExplainCode(ctx context.Context, filetype, code string) (string, error) {
	model := c.gmClient.GenerativeModel("gemini-pro")
	cs := model.StartChat()

	message := fmt.Sprintf("Could you explain me this %s code?\n\n%s", filetype, code)
	res, err := cs.SendMessage(ctx, genai.Text(message))
	if err != nil {
		return "", err
	}
	return c.extractAnswer(res), nil
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

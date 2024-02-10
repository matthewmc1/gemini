package finn

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"

	"cloud.google.com/go/vertexai/genai"
)

type Config struct {
	Project string
	Region  string
}

type Gemini struct {
	Client *genai.Client
	Model  string
}

func NewConfig(context context.Context, c Config, model string) (*Gemini, error) {
	c.Project = os.Getenv("PROJECT_ID")
	c.Region = os.Getenv("REGION")

	client, err := genai.NewClient(context, c.Project, c.Region)
	if err != nil {
		slog.Error("error generating new generative ai client", "client details", client)
	}

	return &Gemini{
		Client: client,
		Model:  model,
	}, nil
}

func (gemini *Gemini) Chat(w io.Writer, context context.Context, request string) {
	model := gemini.Client.GenerativeModel(gemini.Model)

	prompt := genai.Text(request)

	slog.Info(string(prompt))

	model.StartChat()
	response, err := model.StartChat().SendMessage(context, prompt)

	if err != nil {
		slog.Error("error generating response")
		return
	}

	fmt.Fprintf(w, "generated response: %s\n", response.Candidates[0].Content.Parts)
}

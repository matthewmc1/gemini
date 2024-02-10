package main

import (
	"context"
	"finn/finn"
	"log/slog"
	"os"
)

var (
	gemini finn.Gemini
	config finn.Config
)

func main() {
	ctx := context.Background()

	model, err := finn.NewConfig(ctx, config, "gemini-pro")
	slog.Info(model.Model, "Client", model.Client)
	defer model.Client.Close()
	if err != nil {
		slog.Error("error creating configuration")
	}

	text := "Generate a learning plan for a new SRE, you should review https://sre.google/ and other sites including the books to affirm what is important as an SRE and what tooling. Once defined it would be useful to propose some good projects for new SRE's to start working on."

	model.Chat(os.Stdout, ctx, text)
}

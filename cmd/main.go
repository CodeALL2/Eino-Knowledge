package main

import (
	"context"
	"log"

	"Eino-Knowledge/internal/config"
	"Eino-Knowledge/internal/embedding"
	"Eino-Knowledge/internal/llm"
)

func init() {
	config.LoadConfig()
	ctx := context.Background()
	llm.NewLLMModel(ctx)
	embedding.NewEmbedder(ctx)
}

func main() {
	log.Println("Hello World")
}

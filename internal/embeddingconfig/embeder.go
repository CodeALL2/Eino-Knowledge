package embeddingconfig

import (
	"Eino-Knowledge/internal/config"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

type Embedder struct {
	emb *ark.Embedder
}

var embeder *ark.Embedder

func NewEmbedder(ctx context.Context) *Embedder {
	if embeder == nil {
		embConfig := config.GetEmbeddingConfig()
		emb, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
			APIKey:  embConfig.GetAPIKey(),
			BaseURL: embConfig.GetAPIURL(),
			Model:   embConfig.GetModelName(),
		})
		if err != nil {
			log.Println(err)
			return nil
		}
		embeder = emb
	}

	log.Println("嵌入初始化成功")
	return &Embedder{emb: embeder}
}

func GetEmbeder() *ark.Embedder {
	return embeder
}

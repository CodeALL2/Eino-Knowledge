package embedding

import (
	"Eino-Knowledge/internal/config"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
)

type Embedder struct {
	emb *ark.Embedder
}

func NewEmbedder(ctx context.Context) *Embedder {
	emb, err := ark.NewEmbedder(ctx, &ark.EmbeddingConfig{
		APIKey:  config.GetEmbeddingConfig().APIKey,
		BaseURL: config.GetEmbeddingConfig().APIURL,
		Model:   config.GetEmbeddingConfig().ModelName,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("嵌入初始化成功")
	return &Embedder{emb: emb}
}

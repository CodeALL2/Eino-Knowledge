package llm

import (
	"Eino-Knowledge/internal/config"
	"context"
	"log"

	"github.com/cloudwego/eino-ext/components/model/ark"
)

type VllmModel struct {
	model *ark.ChatModel
}

func NewLLMModel(ctx context.Context) *VllmModel {
	model, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey: config.GetModelConfig().APIKey,
		Model:  config.GetModelConfig().ModelName,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("模型初始化成功")
	return &VllmModel{model: model}
}

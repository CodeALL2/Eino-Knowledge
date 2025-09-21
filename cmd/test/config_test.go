package test

import (
	"testing"

	"Eino-Knowledge/internal/config"
)

func TestLoadConfig(t *testing.T) {
	// 测试配置加载
	config.LoadConfig()

	// 验证配置是否正确加载
	embConfig := config.GetEmbeddingConfig()
	if embConfig.APIKey == "" {
		t.Error("嵌入配置的APIKey为空")
	}
	if embConfig.APIURL == "" {
		t.Error("嵌入配置的APIURL为空")
	}
	if embConfig.ModelName == "" {
		t.Error("嵌入配置的ModelName为空")
	}

	modelConfig := config.GetModelConfig()
	if modelConfig.APIKey == "" {
		t.Error("模型配置的APIKey为空")
	}
	if modelConfig.ModelName == "" {
		t.Error("模型配置的ModelName为空")
	}

	t.Log("配置加载测试通过")
}

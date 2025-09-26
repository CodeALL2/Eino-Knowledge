package test

import (
	"context"
	"log"
	"strings"
	"testing"

	"Eino-Knowledge/internal/config"
	"Eino-Knowledge/internal/fileprocessor/loaderengine/loaderimp"

	"github.com/cloudwego/eino/components/document/parser"
)

func TestLoadConfig(t *testing.T) {
	// 测试配置加载
	config.LoadConfig()

	// 验证配置是否正确加载
	embConfig := config.GetEmbeddingConfig()
	if embConfig.GetAPIKey() == "" {
		t.Error("嵌入配置的APIKey为空")
	}
	if embConfig.GetAPIURL() == "" {
		t.Error("嵌入配置的APIURL为空")
	}
	if embConfig.GetModelName() == "" {
		t.Error("嵌入配置的ModelName为空")
	}

	modelConfig := config.GetModelConfig()
	if modelConfig.GetAPIKey() == "" {
		t.Error("模型配置的APIKey为空")
	}
	if modelConfig.GetModelName() == "" {
		t.Error("模型配置的ModelName为空")
	}

	milvusConfig := config.GetMilvusConfig()
	if milvusConfig.GetHost() == "" {
		t.Error("Milvus配置的Host为空")
	}
	if milvusConfig.GetPort() == 0 {
		t.Error("Milvus配置的Port为空")
	}
	if milvusConfig.GetCollection() == "" {
		t.Error("Milvus配置的Collection为空")
	}

	// 测试统一配置接口
	cfg := config.GetConfig()
	if cfg == nil {
		t.Error("统一配置接口返回nil")
	}
	if cfg.Embedding().GetAPIKey() == "" {
		t.Error("统一配置接口中嵌入配置的APIKey为空")
	}

	t.Log("配置加载测试通过")
}

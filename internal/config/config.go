package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var globalConfig *config

type embeddingConfig struct {
	APIKey    string `yaml:"APIKey"`
	APIURL    string `yaml:"APIURL"`
	ModelName string `yaml:"ModelName"`
}

type modelConfig struct {
	APIKey    string `yaml:"APIKey"`
	ModelName string `yaml:"ModelName"`
}

type config struct {
	EmbeddingConfig embeddingConfig `yaml:"Emb"`
	ModelConfig     modelConfig     `yaml:"model"`
}

// findConfigFile 查找配置文件
func findConfigFile() string {
	// 可能的配置文件路径
	possiblePaths := []string{
		"internal/config/modelconfig.yaml",          // 从项目根目录
		"config/modelconfig.yaml",                   // 备选路径
		"modelconfig.yaml",                          // 当前目录
		"./internal/config/modelconfig.yaml",        // 当前目录下
		"../internal/config/modelconfig.yaml",       // 上级目录
		"../../internal/config/modelconfig.yaml",    // 上上级目录
		"../../../internal/config/modelconfig.yaml", // 上上上级目录
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			log.Printf("找到配置文件: %s", path)
			return path
		}
	}

	return ""
}

// LoadConfig 读取配置文件并装填到结构体
func LoadConfig() {
	// 查找配置文件
	configPath := findConfigFile()
	if configPath == "" {
		log.Fatalf("未找到配置文件 modelconfig.yaml")
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}

	// 初始化GlobalConfig
	globalConfig = &config{}

	// 解析YAML到结构体
	err = yaml.Unmarshal(data, globalConfig)
	if err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}

	log.Printf("配置文件加载成功: %s", configPath)
}

// GetEmbeddingConfig 获取嵌入配置
func GetEmbeddingConfig() embeddingConfig {
	if globalConfig != nil {
		return globalConfig.EmbeddingConfig
	}
	return embeddingConfig{}
}

// GetModelConfig 获取模型配置
func GetModelConfig() modelConfig {
	if globalConfig != nil {
		return globalConfig.ModelConfig
	}
	return modelConfig{}
}

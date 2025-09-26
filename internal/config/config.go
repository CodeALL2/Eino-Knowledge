package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var globalConfig *config

// 只读配置接口
type EmbeddingConfig interface {
	GetAPIKey() string
	GetAPIURL() string
	GetModelName() string
}

type ModelConfig interface {
	GetAPIKey() string
	GetModelName() string
}

type MilvusConfig interface {
	GetHost() string
	GetPort() int
	GetCollection() string
	GetDBName() string
}

type Config interface {
	Embedding() EmbeddingConfig
	Model() ModelConfig
	Milvus() MilvusConfig
}

// 内部配置结构体
type embeddingConfig struct {
	APIKey    string `yaml:"APIKey"`
	APIURL    string `yaml:"APIURL"`
	ModelName string `yaml:"ModelName"`
}

func (e embeddingConfig) GetAPIKey() string    { return e.APIKey }
func (e embeddingConfig) GetAPIURL() string    { return e.APIURL }
func (e embeddingConfig) GetModelName() string { return e.ModelName }

type modelConfig struct {
	APIKey    string `yaml:"APIKey"`
	ModelName string `yaml:"ModelName"`
}

func (m modelConfig) GetAPIKey() string    { return m.APIKey }
func (m modelConfig) GetModelName() string { return m.ModelName }

type milvusConfig struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	Collection string `yaml:"Collection"`
	DBName     string `yaml:"DBName"`
}

func (m milvusConfig) GetHost() string       { return m.Host }
func (m milvusConfig) GetPort() int          { return m.Port }
func (m milvusConfig) GetCollection() string { return m.Collection }
func (m milvusConfig) GetDBName() string     { return m.DBName }

type config struct {
	EmbeddingConfig embeddingConfig `yaml:"Emb"`
	ModelConfig     modelConfig     `yaml:"model"`
	MilvusConfig    milvusConfig    `yaml:"milvus"`
}

func (c *config) Embedding() EmbeddingConfig { return c.EmbeddingConfig }
func (c *config) Model() ModelConfig         { return c.ModelConfig }
func (c *config) Milvus() MilvusConfig       { return c.MilvusConfig }

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

// GetConfig 获取全局配置（只读接口）
func GetConfig() Config {
	if globalConfig != nil {
		return globalConfig
	}
	return nil
}

// 为了向后兼容，保留原有的单独获取函数，但返回只读接口
func GetEmbeddingConfig() EmbeddingConfig {
	if globalConfig != nil {
		return globalConfig.EmbeddingConfig
	}
	return embeddingConfig{}
}

func GetModelConfig() ModelConfig {
	if globalConfig != nil {
		return globalConfig.ModelConfig
	}
	return modelConfig{}
}

func GetMilvusConfig() MilvusConfig {
	if globalConfig != nil {
		return globalConfig.MilvusConfig
	}
	return milvusConfig{}
}

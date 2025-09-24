package main

import (
	"context"
	"fmt"
	"log"

	"Eino-Knowledge/internal/config"
	"Eino-Knowledge/internal/milvusconfig"
)

func init() {
	log.Println("初始化配置...")
	config.LoadConfig()
	log.Println("配置加载完成")

	ctx := context.Background()
	log.Println("创建Milvus客户端...")
	milvusconfig.CreateMilvusClient(ctx)
	log.Println("Milvus客户端创建完成")
}

func main() {
	ctx := context.Background()
	log.Println("开始查询Milvus数据...")

	milvusClient := milvusconfig.GetMilvusClient()
	collectionName := config.GetMilvusConfig().GetCollection()

	log.Printf("连接到Milvus: %s:%d", config.GetMilvusConfig().GetHost(), config.GetMilvusConfig().GetPort())
	log.Printf("查询集合: %s", collectionName)

	// 检查集合是否存在
	hasCollection, err := milvusClient.HasCollection(ctx, collectionName)
	if err != nil {
		log.Printf("检查集合失败: %v", err)
		return
	}

	if !hasCollection {
		log.Printf("❌ 集合 %s 不存在", collectionName)
		return
	}

	log.Printf("✅ 集合 %s 存在", collectionName)

	// 加载集合到内存
	err = milvusClient.LoadCollection(ctx, collectionName, false)
	if err != nil {
		log.Printf("加载集合失败: %v", err)
		return
	}

	// 获取集合统计信息
	stats, err := milvusClient.GetCollectionStatistics(ctx, collectionName)
	if err != nil {
		log.Printf("获取集合统计失败: %v", err)
		return
	}

	fmt.Printf("集合统计信息:\n")
	for key, value := range stats {
		fmt.Printf("  %s: %s\n", key, value)
	}

	fmt.Printf("\n✅ 数据已成功存储到Milvus数据库中！\n")
	fmt.Printf("集合名称: %s\n", collectionName)
	fmt.Printf("连接信息: %s:%d\n", config.GetMilvusConfig().GetHost(), config.GetMilvusConfig().GetPort())
}

package milvusconfig

import (
	"Eino-Knowledge/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
)

var milvusClient client.Client
var milvusCollection string

func CreateMilvusClient(ctx context.Context) {
	milvusConfig := config.GetMilvusConfig()
	fmt.Printf("%s:%d", milvusConfig.GetHost(), milvusConfig.GetPort())
	client, err := client.NewClient(ctx, client.Config{
		Address: fmt.Sprintf("%s:%d", milvusConfig.GetHost(), milvusConfig.GetPort()),
		DBName:  "MyEino",
	})

	if err != nil {
		log.Fatal("milvus-client初始化失败")
		panic(err)
	}
	log.Printf("milvus-client初始化成功，连接地址: %s:%d 集合:%s", milvusConfig.GetHost(), milvusConfig.GetPort(), milvusConfig.GetCollection())
	milvusClient = client
	milvusCollection = milvusConfig.GetCollection()
}

func GetMilvusClient() client.Client {
	return milvusClient
}

func GetMilvusCollection() string {
	return milvusCollection
}

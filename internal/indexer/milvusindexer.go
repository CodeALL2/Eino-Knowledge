package indexer

import (
	"context"
	"log"

	"Eino-Knowledge/internal/embeddingconfig"
	"Eino-Knowledge/internal/indexer/milvusfiled"
	"Eino-Knowledge/internal/milvusconfig"

	"github.com/cloudwego/eino-ext/components/indexer/milvus"
	"github.com/cloudwego/eino/schema"
)

type MiluvsIndexer struct {
	indexer *milvus.Indexer
}

func NewIndexer(ctx context.Context) *MiluvsIndexer {
	index, err := milvus.NewIndexer(ctx, &milvus.IndexerConfig{
		Client:     milvusconfig.GetMilvusClient(),     //milvus客户端
		Collection: milvusconfig.GetMilvusCollection(), //collection名字
		Fields:     milvusfiled.MilvusFiled,
		Embedding:  embeddingconfig.GetEmbeder(),
	})

	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("milvus索引初始化成功", milvusconfig.GetMilvusCollection())
	return &MiluvsIndexer{indexer: index}
}

func (m *MiluvsIndexer) StoreBatch(ctx context.Context, docs []*schema.Document) error {
	ids, err := m.indexer.Store(ctx, docs)
	if err != nil {
		log.Println(err)
		log.Println("milvus索引存储失败")
		return err
	}
	log.Println("milvus索引存储成功:", ids)
	return nil
}

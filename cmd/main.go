package main

import (
	"context"
	"log"

	"Eino-Knowledge/internal/config"
	"Eino-Knowledge/internal/embeddingconfig"
	"Eino-Knowledge/internal/indexer"
	"Eino-Knowledge/internal/llm"
	"Eino-Knowledge/internal/milvusconfig"

	"github.com/cloudwego/eino/schema"
	"github.com/google/uuid"
)

func init() {
	config.LoadConfig()
	ctx := context.Background()
	llm.NewLLMModel(ctx)
	embeddingconfig.NewEmbedder(ctx)
	milvusconfig.CreateMilvusClient(ctx)
}

func main() {
	ctx := context.Background()

	// 初始化索引器
	milvusIndexer := indexer.NewIndexer(ctx)
	if milvusIndexer == nil {
		log.Fatal("索引器初始化失败")
		return
	}

	// 创建关于Redis的文档数据
	redisDocuments := createRedisDocuments()

	// 存储文档到Milvus
	err := milvusIndexer.StoreBatch(ctx, redisDocuments)
	if err != nil {
		log.Printf("文档存储失败: %v", err)
		return
	}

	log.Println("Redis知识库文档存储完成!")
}

// createRedisDocuments 创建关于Redis的文档
func createRedisDocuments() []*schema.Document {
	documents := []*schema.Document{
		{
			ID: "1",
			Content: `Redis是什么？
Redis（Remote Dictionary Server）是一个开源的内存数据结构存储系统，可以用作数据库、缓存和消息代理。它支持多种数据结构，如字符串、哈希、列表、集合、有序集合等。Redis以其高性能、低延迟和丰富的功能而闻名，广泛用于Web应用程序、实时分析和缓存层。`,
			MetaData: map[string]interface{}{
				"title":    "Redis简介",
				"category": "基础概念",
				"tags":     []string{"Redis", "NoSQL", "内存数据库", "缓存"},
				"source":   "知识库",
			},
		},
		{
			ID: uuid.New().String(),
			Content: `Redis的核心数据类型
1. String（字符串）：最基本的数据类型，可以存储文本、数字或二进制数据
2. Hash（哈希）：键值对的集合，类似于对象或字典
3. List（列表）：有序的字符串列表，支持头尾操作
4. Set（集合）：无序的唯一字符串集合
5. Sorted Set（有序集合）：带分数的有序集合
6. Bitmap（位图）：二进制位操作
7. HyperLogLog：基数统计
8. Stream：日志数据结构`,
			MetaData: map[string]interface{}{
				"title":    "Redis数据类型",
				"category": "数据结构",
				"tags":     []string{"Redis", "数据类型", "String", "Hash", "List", "Set"},
				"source":   "知识库",
			},
		},
		{
			ID: uuid.New().String(),
			Content: `Redis持久化机制
Redis提供两种持久化方式：
1. RDB（Redis Database）：
   - 在指定时间间隔内生成数据集的快照
   - 适合备份和灾难恢复
   - 文件紧凑，恢复速度快
   
2. AOF（Append Only File）：
   - 记录每个写操作的日志
   - 数据安全性更高
   - 文件较大，恢复速度相对较慢
   
可以同时启用RDB和AOF，Redis会优先使用AOF恢复数据。`,
			MetaData: map[string]interface{}{
				"title":    "Redis持久化",
				"category": "数据持久化",
				"tags":     []string{"Redis", "RDB", "AOF", "持久化", "备份"},
				"source":   "知识库",
			},
		},
		{
			ID: uuid.New().String(),
			Content: `Redis集群和高可用性
1. 主从复制（Master-Slave）：
   - 主节点处理写操作，从节点处理读操作
   - 提供读写分离和数据冗余
   
2. Redis Sentinel：
   - 监控主从节点的健康状态
   - 自动故障转移
   - 配置管理和通知
   
3. Redis Cluster：
   - 数据分片存储
   - 高可用和可扩展性
   - 无中心架构`,
			MetaData: map[string]interface{}{
				"title":    "Redis集群",
				"category": "高可用性",
				"tags":     []string{"Redis", "集群", "主从复制", "Sentinel", "高可用"},
				"source":   "知识库",
			},
		},
		{
			ID: uuid.New().String(),
			Content: `Redis常用命令
字符串操作：
- SET key value：设置键值
- GET key：获取值
- INCR key：数值递增
- DECR key：数值递减

哈希操作：
- HSET key field value：设置哈希字段
- HGET key field：获取哈希字段值
- HGETALL key：获取所有字段和值

列表操作：
- LPUSH key value：左侧插入
- RPUSH key value：右侧插入
- LPOP key：左侧弹出
- RPOP key：右侧弹出

集合操作：
- SADD key member：添加成员
- SMEMBERS key：获取所有成员
- SINTER key1 key2：交集运算`,
			MetaData: map[string]interface{}{
				"title":    "Redis常用命令",
				"category": "操作命令",
				"tags":     []string{"Redis", "命令", "SET", "GET", "HSET", "LPUSH"},
				"source":   "知识库",
			},
		},
		{
			ID: uuid.New().String(),
			Content: `Redis性能优化
1. 内存优化：
   - 选择合适的数据结构
   - 设置过期时间
   - 使用内存分析工具
   
2. 网络优化：
   - 使用管道（Pipeline）批量操作
   - 减少网络往返次数
   - 使用连接池
   
3. CPU优化：
   - 避免耗时操作
   - 使用合适的数据类型
   - 优化Lua脚本
   
4. 持久化优化：
   - 合理配置RDB和AOF
   - 使用SSD存储
   - 调整写入策略`,
			MetaData: map[string]interface{}{
				"title":    "Redis性能优化",
				"category": "性能调优",
				"tags":     []string{"Redis", "性能优化", "内存", "网络", "CPU", "持久化"},
				"source":   "知识库",
			},
		},
		{
			ID: uuid.New().String(),
			Content: `Redis应用场景
1. 缓存系统：
   - Web应用缓存
   - 数据库查询缓存
   - 会话存储
   
2. 实时计数器：
   - 网站访问统计
   - 点赞数、评论数
   - 库存计数
   
3. 消息队列：
   - 任务队列
   - 发布订阅
   - 延时队列
   
4. 实时排行榜：
   - 游戏排行榜
   - 热门内容排序
   - 用户积分排名
   
5. 分布式锁：
   - 防止重复操作
   - 资源互斥访问`,
			MetaData: map[string]interface{}{
				"title":    "Redis应用场景",
				"category": "实际应用",
				"tags":     []string{"Redis", "缓存", "计数器", "消息队列", "排行榜", "分布式锁"},
				"source":   "知识库",
			},
		},
	}

	return documents
}

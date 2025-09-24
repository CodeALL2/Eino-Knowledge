package milvusfiled

import "github.com/milvus-io/milvus-sdk-go/v2/entity"

var MilvusFiled = []*entity.Field{
	{
		Name:       "id",
		DataType:   entity.FieldTypeVarChar,
		PrimaryKey: true,
		TypeParams: map[string]string{
			"max_length": "256",
		},
	},
	{
		Name:     "vector",
		DataType: entity.FieldTypeBinaryVector,
		TypeParams: map[string]string{
			"dim": "81920",
		},
	},
	{
		Name:     "content",
		DataType: entity.FieldTypeVarChar,
		TypeParams: map[string]string{
			"max_length": "8192",
		},
	},
	{
		Name:     "metadata",
		DataType: entity.FieldTypeJSON,
	},
}

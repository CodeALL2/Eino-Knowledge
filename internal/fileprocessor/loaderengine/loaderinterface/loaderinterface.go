package loaderinterface

import (
	"context"

	"github.com/cloudwego/eino/schema"
)

type LoaderInterface interface {
	Load(ctx context.Context, filePath string, metadata map[string]any) ([]*schema.Document, error)
}

package loaderimp

import (
	"context"

	"github.com/cloudwego/eino-ext/components/document/loader/file"
	"github.com/cloudwego/eino-ext/components/document/parser/html"
	"github.com/cloudwego/eino-ext/components/document/parser/pdf"
	"github.com/cloudwego/eino/components/document"
	"github.com/cloudwego/eino/components/document/parser"
	"github.com/cloudwego/eino/schema"
)

type LoaderEngine struct {
}

func NewLoaderEngine(ctx context.Context) *LoaderEngine {
	return &LoaderEngine{}
}

/*
*
加载文件的入口
*/
func (l *LoaderEngine) Load(ctx context.Context, filePath string, metadata map[string]any) ([]*schema.Document, error) {
	textParser := &parser.TextParser{} //默认的文本类加载器
	pdfParser, err := pdf.NewPDFParser(ctx, &pdf.Config{ToPages: false})
	htmlParser, err := html.NewParser(ctx, &html.Config{Selector: ptrString("body")})
	if err != nil {
		panic(err)
	}
	extParser, err := parser.NewExtParser(ctx, &parser.ExtParserConfig{
		FallbackParser: textParser,
		Parsers: map[string]parser.Parser{
			".pdf":  pdfParser,
			".html": htmlParser,
		},
	})
	if err != nil {
		panic(err)
	}
	fileLoader, err := file.NewFileLoader(ctx, &file.FileLoaderConfig{
		Parser: extParser,
	})
	if err != nil {
		panic(err)
	}
	return fileLoader.Load(ctx, document.Source{URI: filePath})
}

func ptrString(s string) *string {
	return &s
}

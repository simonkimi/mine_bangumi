package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/graph/grmodel"
	"github.com/simonkimi/minebangumi/internal/app/service/source_parser"
)

type Resolver struct{}

func (r *Resolver) ParseSource(ctx context.Context, source *grmodel.ParseSourceInput) (*grmodel.ParseSourceResponse, error) {
	result, err := source_parser.ParseSource(ctx, source.Source, source.Parser)
	if err != nil {
		return nil, err
	}

	var files []*string
	for _, file := range result.Files {
		files = append(files, &file)
	}

	return &grmodel.ParseSourceResponse{
		Title:  &result.RawTitle,
		Files:  files,
		Season: &result.Season,
	}, nil
}

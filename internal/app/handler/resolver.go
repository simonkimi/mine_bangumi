package handler

import (
	"context"
	"github.com/simonkimi/minebangumi/internal/app/graph/grmodel"
	"github.com/simonkimi/minebangumi/internal/app/service/source_parser"
)

type Resolver struct{}

func (r *Resolver) ParseSource(ctx context.Context, input *grmodel.ParseSourceInput) (*grmodel.ParseSourceResponse, error) {
	source, err := source_parser.ParseSource(ctx, input.Source, input.Parser)
	if err != nil {
		return nil, err
	}

	var files []*string
	for _, file := range source.Files {
		files = append(files, &file)
	}

	return &grmodel.ParseSourceResponse{
		Title:  &source.RawTitle,
		Files:  files,
		Season: &source.Season,
	}, nil
}

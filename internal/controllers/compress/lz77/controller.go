package lz77

import (
	"context"

	"course-compress/pkg/model"
)

type CompressorService interface {
	Encode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error)
	Decode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error)
}

type Config struct {
	CompressorService CompressorService
}

type Handler struct {
	compressorService CompressorService
}

func NewHandler(cfg Config) *Handler {
	return &Handler{
		compressorService: cfg.CompressorService,
	}
}

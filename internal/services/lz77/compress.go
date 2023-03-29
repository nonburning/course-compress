package lz77

import (
	"context"
	"fmt"

	"course-compress/pkg/model"
)

type Processor interface {
	Lz77Encode(ctx context.Context, bytes string) (*string, error)
	Lz77Decode(ctx context.Context, bytes string) (*string, error)
}

type CompressService struct {
	processor Processor
}

func (c *CompressService) Encode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error) {
	bytes, err := c.processor.Lz77Encode(ctx, pointer.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encode data with LZ77 algorithm: %v", err)
	}
	return &model.Pointer{
		Bytes: *bytes,
	}, nil
}

func (c *CompressService) Decode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error) {
	bytes, err := c.processor.Lz77Decode(ctx, pointer.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data with LZ77 algorithm: %v", err)
	}
	return &model.Pointer{
		Bytes: *bytes,
	}, nil
}

func NewCompressService(processor Processor) *CompressService {
	return &CompressService{processor: processor}
}

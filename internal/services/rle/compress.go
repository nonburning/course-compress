package rle

import (
	"context"
	"fmt"

	"course-compress/pkg/model"
)

type Processor interface {
	RleEncode(ctx context.Context, bytes string) (*string, error)
	RleDecode(ctx context.Context, bytes string) (*string, error)
}

type CompressService struct {
	processor Processor
}

func (c *CompressService) Encode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error) {
	bytes, err := c.processor.RleEncode(ctx, pointer.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encode data with REL algorithm: %v", err)
	}
	return &model.Pointer{
		Bytes: *bytes,
	}, nil
}

func (c *CompressService) Decode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error) {
	bytes, err := c.processor.RleDecode(ctx, pointer.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data with REL algorithm: %v", err)
	}
	return &model.Pointer{
		Bytes: *bytes,
	}, nil
}

func NewCompressService(processor Processor) *CompressService {
	return &CompressService{processor: processor}
}

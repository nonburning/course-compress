package bwt

import (
	"context"
	"fmt"

	"course-compress/pkg/model"
)

type Processor interface {
	BwtEncode(ctx context.Context, bytes string) (*string, error)
	BwtDecode(ctx context.Context, bytes string) (*string, error)
}

type CompressService struct {
	processor Processor
}

func (c *CompressService) Encode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error) {
	bytes, err := c.processor.BwtEncode(ctx, pointer.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to encode data with BWT algorithm: %v", err)
	}
	return &model.Pointer{
		Bytes: *bytes,
	}, nil
}

func (c *CompressService) Decode(ctx context.Context, pointer *model.Pointer) (*model.Pointer, error) {
	bytes, err := c.processor.BwtDecode(ctx, pointer.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode data with BWT algorithm: %v", err)
	}
	return &model.Pointer{
		Bytes: *bytes,
	}, nil
}

func NewCompressService(processor Processor) *CompressService {
	return &CompressService{processor: processor}
}

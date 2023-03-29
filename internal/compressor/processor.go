package compressor

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "library.h"
import "C"
import (
	"context"
	"encoding/base64"
	"fmt"
	"unsafe"
)

type Processor struct{}

func NewProcessor() *Processor { return &Processor{} }

func newStructPtr(bytes string) (*C.struct_ptr, error) {
	origin, err := base64.StdEncoding.DecodeString(bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 string: %v", err)
	}
	return &C.struct_ptr{
		bytes:  (*C.uchar)(C.CBytes(origin)),
		length: C.uint(len(origin)),
	}, nil
}

func (r *Processor) RleEncode(ctx context.Context, bytes string) (*string, error) {
	input, err := newStructPtr(bytes)
	if err != nil {
		return nil, err
	}

	var encoded *C.struct_ptr = C.rle_encode(input)
	body := base64.StdEncoding.EncodeToString(C.GoBytes(unsafe.Pointer(encoded.bytes), C.int(encoded.length)))
	return &body, nil
}

func (r *Processor) RleDecode(ctx context.Context, bytes string) (*string, error) {
	input, err := newStructPtr(bytes)
	if err != nil {
		return nil, err
	}

	var decoded *C.struct_ptr = C.rle_decode(input)
	body := base64.StdEncoding.EncodeToString(C.GoBytes(unsafe.Pointer(decoded.bytes), C.int(decoded.length)))
	return &body, nil
}

func (r *Processor) Lz77Encode(ctx context.Context, bytes string) (*string, error) {
	input, err := newStructPtr(bytes)
	if err != nil {
		return nil, err
	}

	var encoded *C.struct_ptr = C.lz77_encode(input)
	body := base64.StdEncoding.EncodeToString(C.GoBytes(unsafe.Pointer(encoded.bytes), C.int(encoded.length)))
	return &body, nil
}

func (r *Processor) Lz77Decode(ctx context.Context, bytes string) (*string, error) {
	input, err := newStructPtr(bytes)
	if err != nil {
		return nil, err
	}

	var encoded *C.struct_ptr = C.lz77_decode(input)
	body := base64.StdEncoding.EncodeToString(C.GoBytes(unsafe.Pointer(encoded.bytes), C.int(encoded.length)))
	return &body, nil
}

func (r *Processor) BwtEncode(ctx context.Context, bytes string) (*string, error) {
	input, err := newStructPtr(bytes)
	if err != nil {
		return nil, err
	}

	var encoded *C.struct_ptr = C.bwt_encode(input)
	body := base64.StdEncoding.EncodeToString(C.GoBytes(unsafe.Pointer(encoded.bytes), C.int(encoded.length)))
	return &body, nil
}

func (r *Processor) BwtDecode(ctx context.Context, bytes string) (*string, error) {
	input, err := newStructPtr(bytes)
	if err != nil {
		return nil, err
	}

	var encoded *C.struct_ptr = C.bwt_decode(input)
	body := base64.StdEncoding.EncodeToString(C.GoBytes(unsafe.Pointer(encoded.bytes), C.int(encoded.length)))
	return &body, nil
}

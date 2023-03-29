package dto

import "course-compress/pkg/model"

type ProcessPointerReq struct {
	Bytes string `json:"bytes"`
}

func (p *ProcessPointerReq) Pointer() *model.Pointer {
	return &model.Pointer{Bytes: p.Bytes}
}

type ProcessPointerRes struct {
	Bytes string `json:"bytes"`
}

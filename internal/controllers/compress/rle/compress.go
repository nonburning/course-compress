package rle

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/exp/slog"

	"course-compress/internal/dto"
)

func (h *Handler) Encode(e echo.Context) error {
	slog.Info("processing https request", slog.Attr{Key: "method", Value: slog.StringValue(e.Request().Method)}, slog.Attr{Key: "path", Value: slog.StringValue(e.Path())})

	var req dto.ProcessPointerReq
	if err := e.Bind(&req); err != nil {
		slog.Error("failed to bind req body", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		return echo.NewHTTPError(http.StatusBadRequest, dto.HTTPMessage{Message: "invalid request body"})
	}
	ptr := req.Pointer()

	res, err := h.compressorService.Encode(context.Background(), ptr)
	if err != nil {
		slog.Error("error occurred", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		return echo.NewHTTPError(http.StatusInternalServerError, dto.HTTPMessage{Message: "internal server error"})
	}
	return e.JSON(http.StatusOK, dto.ProcessPointerRes{Bytes: res.Bytes})
}

func (h *Handler) Decode(e echo.Context) error {
	slog.Info("processing https request", slog.Attr{Key: "method", Value: slog.StringValue(e.Request().Method)}, slog.Attr{Key: "path", Value: slog.StringValue(e.Path())})

	var req dto.ProcessPointerReq
	if err := e.Bind(&req); err != nil {
		slog.Error("failed to bind req body", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		return echo.NewHTTPError(http.StatusBadRequest, dto.HTTPMessage{Message: "invalid request body"})
	}
	ptr := req.Pointer()

	res, err := h.compressorService.Decode(context.Background(), ptr)
	if err != nil {
		slog.Error("error occurred", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		return echo.NewHTTPError(http.StatusInternalServerError, dto.HTTPMessage{Message: "internal server error"})
	}
	return e.JSON(http.StatusOK, dto.ProcessPointerRes{Bytes: res.Bytes})
}

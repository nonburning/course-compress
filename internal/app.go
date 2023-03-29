package internal

import (
	bwtsvc "course-compress/internal/services/bwt"
	"net/http"

	"github.com/labstack/echo"

	"course-compress/internal/compressor"
	bwthndlr "course-compress/internal/controllers/compress/bwt"
	lz77hndlr "course-compress/internal/controllers/compress/lz77"
	rlehndlr "course-compress/internal/controllers/compress/rle"
	lz77svc "course-compress/internal/services/lz77"
	rlesvc "course-compress/internal/services/rle"
)

type routerConfig struct {
	e *echo.Echo

	relHandler  *rlehndlr.Handler
	lz77Handler *lz77hndlr.Handler
	bwtHandler  *bwthndlr.Handler
}

func newRouterConfig(e *echo.Echo) *routerConfig {
	processor := compressor.NewProcessor()

	return &routerConfig{
		e:           e,
		relHandler:  rlehndlr.NewHandler(rlehndlr.Config{CompressorService: rlesvc.NewCompressService(processor)}),
		lz77Handler: lz77hndlr.NewHandler(lz77hndlr.Config{CompressorService: lz77svc.NewCompressService(processor)}),
		bwtHandler:  bwthndlr.NewHandler(bwthndlr.Config{CompressorService: bwtsvc.NewCompressService(processor)}),
	}
}

func NewApplication() http.Handler {
	handler := setupHandler()
	routerCfg := newRouterConfig(handler)
	setupRoutes(routerCfg)
	return handler
}

func setupHandler() *echo.Echo {
	e := echo.New()
	return e
}

func setupRoutes(cfg *routerConfig) {
	base := cfg.e

	relApi := base.Group("/rle")
	{
		relApi.POST("/encode", cfg.relHandler.Encode)
		relApi.POST("/decode", cfg.relHandler.Decode)
	}
	lz77Api := base.Group("/lz77")
	{
		lz77Api.POST("/encode", cfg.lz77Handler.Encode)
		lz77Api.POST("/decode", cfg.lz77Handler.Decode)
	}
	bwtApi := base.Group("/bwt")
	{
		bwtApi.POST("/encode", cfg.bwtHandler.Encode)
		bwtApi.POST("/decode", cfg.bwtHandler.Decode)
	}
}

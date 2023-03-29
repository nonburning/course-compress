package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"

	"course-compress/internal"
	loggerpkg "course-compress/pkg/logger"
)

const BindPort = "8080"

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	setupLogger()

	handler := internal.NewApplication()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", BindPort),
		Handler: logResponse(handler),
	}
	gr, gtCtx := errgroup.WithContext(ctx)
	gr.Go(serve(srv, gtCtx))
	if err := gr.Wait(); err != nil {
		slog.Error("failed to run application", slog.Attr{
			Key:   "error",
			Value: slog.StringValue(err.Error()),
		})
	}
}

func setupLogger() {
	opts := loggerpkg.PrettyHandlerOptions{SlogOpts: slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}}
	handler := opts.NewPrettyHandler(os.Stdout)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

type statusResponseWriter struct {
	http.ResponseWriter
	statusCode    int
	headerWritten bool
}

func newStatusResponseWriter(w http.ResponseWriter) *statusResponseWriter {
	return &statusResponseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (s *statusResponseWriter) WriteHeader(statusCode int) {
	s.ResponseWriter.WriteHeader(statusCode)

	if !s.headerWritten {
		s.statusCode = statusCode
		s.headerWritten = true
	}
}

func (s *statusResponseWriter) Write(b []byte) (int, error) {
	s.headerWritten = true
	return s.ResponseWriter.Write(b)
}

func (s *statusResponseWriter) Unwrap() http.ResponseWriter {
	return s.ResponseWriter
}

func logResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := newStatusResponseWriter(w)
		next.ServeHTTP(sw, r)
		slog.Info("request processed", slog.Attr{Key: "method", Value: slog.StringValue(r.Method)}, slog.Attr{Key: "path", Value: slog.StringValue(r.URL.Path)}, slog.Attr{Key: "status_code", Value: slog.IntValue(sw.statusCode)})
	})
}

func serve(srv *http.Server, ctx context.Context) func() error {
	return func() error {
		gr, ctx := errgroup.WithContext(ctx)
		gr.Go(func() error {
			slog.Info("API listening", slog.Attr{
				Key:   "port",
				Value: slog.StringValue(BindPort),
			})

			if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		})
		gr.Go(func() error {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 100)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil && !errors.Is(err, http.ErrServerClosed) {
				return err
			}
			return nil
		})
		return gr.Wait()
	}
}

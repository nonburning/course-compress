package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/fatih/color"
	"golang.org/x/exp/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type PrettyHandler struct {
	slog.Handler

	opts PrettyHandlerOptions
	l    *log.Logger
}

func (opts PrettyHandlerOptions) NewPrettyHandler(out io.Writer) *PrettyHandler {
	return &PrettyHandler{
		Handler: opts.SlogOpts.NewJSONHandler(out),
		l:       log.New(out, "", 0),
	}
}

func (p *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(attr slog.Attr) {
		fields[attr.Key] = attr.Value.Any()
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	p.l.Println(timeStr, level, msg, color.WhiteString(string(b)))
	return nil
}

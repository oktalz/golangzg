package main

import (
	"log/slog"
	"os"
	"strings"
)

// START OMIT
func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "source" {
				x := a.Value
				src := x.Any().(*slog.Source)
				path := strings.Split(src.File, "/")
				src.File = path[len(path)-1] // short path
				return slog.Attr{
					Key:   "source",
					Value: slog.AnyValue(src),
				}
			}
			return a
		},
	}))
	slog.SetDefault(logger)

	slog.Info("Hello world!") // HL
}

// END OMIT

// RUN OMIT
// NUR OMIT

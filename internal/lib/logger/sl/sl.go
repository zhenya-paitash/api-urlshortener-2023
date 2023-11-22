package sl

import (
	"fmt"
	"log/slog"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}

// TODO: refactor to generic
func Str(key string, value interface{}) slog.Attr {
  return slog.String(key, fmt.Sprintf("%v", value))
}

func Int64(key string, value int64) slog.Attr {
  return slog.Int64(key, value)
}

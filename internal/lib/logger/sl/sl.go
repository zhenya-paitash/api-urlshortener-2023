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

func Str(key string, value interface{}) slog.Attr {
  return slog.String(key, fmt.Sprintf("%v", value))
}

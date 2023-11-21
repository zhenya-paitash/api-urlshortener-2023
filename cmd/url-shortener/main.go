package main

import (
	"fmt"

	"github.com/zhenya-paitash/api-urlshortener-2023/internal/config"
)

func main() {
	// NOTE: config: cleanenv
	config := config.MustLoad()
	fmt.Println(config)

	// TODO: logger
	// NOTE: logger: slog

	// TODO: storage
	// NOTE: logger: sqlite

	// TODO: router
	// NOTE: router: chi, chi-render

	// TODO: run server
}

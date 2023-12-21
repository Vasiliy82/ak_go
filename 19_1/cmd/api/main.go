package main

import (
	"context"
	"fmt"
	"lesson/internal/api/app"
	"os"
)

func main() {
	ctx := context.Background()
	app.Start(ctx, fmt.Sprintf(":%s",
		os.Getenv("LESSON_API_HTTP_PORT")), fmt.Sprintf("%s:%s",
		os.Getenv("LESSON_API_POPULARITY_URL"),
		os.Getenv("LESSON_API_POPULARITY_PORT")))
}

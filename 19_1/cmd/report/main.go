package main

import (
	"context"
	"fmt"
	"lesson/internal/report/app"
	"os"
)

func main() {
	ctx := context.Background()
	app.Start(ctx, fmt.Sprintf(":%s",
		os.Getenv("LESSON_REPORT_GRPC_PORT")))
}

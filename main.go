package main

import (
	"context"
	"fitup/config"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	app := &config.App{}
	app.Initialize(ctx)
	app.Run("5000")
}

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/LYH2263/go-fetch-timeout/internal/fetch"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	_, err := fetch.New(nil).Get(ctx, "http://127.0.0.1:1")
	fmt.Println(err)
}

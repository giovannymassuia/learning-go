package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "language", "Go")

	bookHotel(ctx, "My Hotel")
}

// ctx parameters, as per convention, should be the first parameter
func bookHotel(ctx context.Context, name string) {
	lang := ctx.Value("language")
	fmt.Println("bookHotel: language is", lang)
}

package main

import(
	"fmt"
	"context"
)

func main() {
	type strKey string

	f := func(ctx context.Context, k strKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value: ", v)
			return
		}

		fmt.Println("key not found: ", k)
	}

	k := strKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	// 将key设置到了context中
	f(ctx, k)
	// 这个没有
	f(ctx, strKey("coloer"))
}
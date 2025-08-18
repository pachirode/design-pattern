package context

import (
	"context"
	"fmt"
	"time"
)

type RequestInfo struct {
	Username string
}

func processRequest(ctx context.Context, req RequestInfo) {
	username := ctx.Value("request").(RequestInfo).Username
	fmt.Printf("用户: %s 发送请求", username)
	time.Sleep(2 * time.Second)
}

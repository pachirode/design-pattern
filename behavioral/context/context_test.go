package context

import (
	"context"
	"testing"
)

func TestContext(t *testing.T) {
	ctx := context.Background()
	req := RequestInfo{
		Username: "张三",
	}

	ctx = context.WithValue(ctx, "request", req)
	processRequest(ctx, req)
}

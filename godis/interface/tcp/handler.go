package tcp

import (
	"context"
	"net"
)

// HandleFunc represents application handler function
type HandleFunc func(ctx context.Context, conn net.Conn)

// handler 应用层服务的抽象
type Handler interface {
	Handle(ctx context.Context, conn net.Conn)
	Close()error
}


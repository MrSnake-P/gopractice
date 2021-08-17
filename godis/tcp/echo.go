package tcp

import (
	"bufio"
	"context"
	"godis/lib/logger"
	"godis/lib/sync/atomic"
	"godis/lib/sync/wait"
	"io"
	"net"
	"sync"
)

type Client struct {
	Conn net.Conn

	Waiting wait.Wait
}

type EchoHandler struct {
	// 保存所有工作状态client的集合（map当set用）
	// 需要并发安全，map不是并发安全的
	/* 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。 */
	activeConn sync.Map

	// 关闭状态标识位
	closing    atomic.Boolean
}

func (h *EchoHandler) Handler (ctx context.Context, conn net.Conn) {
	// 关闭中的handler不会处理新连接
	if h.closing.Get() {
		conn.Close()
	}

	client := &Client{
		Conn: conn,
	}
	h.activeConn.Store(client, struct{}{}) // 记住仍然存活的连接

	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				logger.Info("connection close")
				h.activeConn.Delete(conn)
			} else {
				logger.Warn(err)
			}
			return
		}
		//

	}
}
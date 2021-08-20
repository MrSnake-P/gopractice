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
	"time"
)

type Client struct {
	Conn net.Conn

	Waiting wait.Wait
}

func MakeEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

type EchoHandler struct {
	// 保存所有工作状态client的集合（map当set用）
	// 需要并发安全，map不是并发安全的
	/* 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。 */
	activeConn sync.Map

	// 关闭状态标识位
	closing    atomic.Boolean
}

func (h *EchoHandler) Handle (ctx context.Context, conn net.Conn) {
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
		// 发送数据先置为waiting状态，阻止连接被关闭
		client.Waiting.Add(1)

		// 模拟关闭时未完成发送的情况
		//logger.Info("sleeping")
		//time.Sleep(10 * time.Second)
		b := []byte(msg)
		conn.Write(b)
		// 发送完毕
		client.Waiting.Done()
	}
}

// 关闭客户端连接
func (c *Client) Close() error {
	// 等待数据发送完成或超时
	c.Waiting.WaitWithTimeOut(10 * time.Second)
	c.Conn.Close()
	return nil
}

// 关闭服务器
func (h *EchoHandler) Close() error {
	logger.Info("handler shtting down...")
	h.closing.Set(true)
	// 逐个关闭
	h.activeConn.Range(func(key interface{}, val interface{})bool{
		client := key.(*Client)
		client.Close()
		return true
	})
	return nil
}
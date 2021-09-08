package main

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"log"
)

type testServer struct {
	*gnet.EventServer
	addr      string
	multicore bool
	async     bool
	pool      *goroutine.Pool
}

func (ts *testServer) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Fatalf("Server connect to %s", ts.addr)
	return
}

func (ts *testServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (ts *testServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	log.Fatalf("Server closed")
	return
}

func (ts *testServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if ts.async {
		data := append([]byte{}, frame...)
		_ = ts.pool.Submit(func() {
			c.AsyncWrite(data)
		})
		return
	}
	out = frame
	return

}

func main() {
	a := fmt.Sprintf("tcp://:%d", 9000)
	test := &testServer{
		addr:      a,
		multicore: true,
		async:     false,
		pool:      goroutine.Default()}
	err := gnet.Serve(test,
		a,
		gnet.WithMulticore(true),
		gnet.WithReusePort(true))
	if err != nil {
		panic(err)
	}
}

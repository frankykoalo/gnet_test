package test

import (
	"context"
	"fmt"
	"gnet_test/proto/proto"
	"google.golang.org/grpc"
	"log"
	"testing"
	"time"
)

func TestGrpc(t *testing.T) {
	conn, _ := grpc.Dial("127.0.0.1:9001", grpc.WithInsecure())
	defer conn.Close()
	c := proto.NewEchoServiceClient(conn)
	i := 97
	b := make([]byte, 0)
	for {
		time.Sleep(1 * time.Second)
		b = append(b, byte(i))
		fmt.Printf("Send data is %s\n", string(b))
		out, err := c.Echo(context.Background(), &proto.Request{Input: b})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive data is %s\n", string(out.Output))
		if i < 122 {
			i++
		} else {
			break
		}
	}
}

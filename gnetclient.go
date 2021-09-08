package encode

import (
	"fmt"
	"github.com/sammyne/base58"
	"net"
	"time"
)

func StartClient(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connecting..")

	i := 97
	b := make([]byte, 0)
	bytes := make([]byte, 1024)
	for {
		b = append(b, byte(i))
		_, err = conn.Write(b)
		fmt.Printf("Send data is %v\n", string(b))
		if err != nil {
			fmt.Printf("%s", err)
			break
		}
		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Printf("%s", err)
			break
		}
		result := base58.CheckEncodeX(bytes)
		fmt.Printf("Receive data is %v\n", result)
		if i < 122 {
			time.Sleep(3 * time.Second)
			i++
		} else {
			break
		}
	}
}
func StartEchoClient(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Printf("dial failed, err:%v\n", err)
		return
	}

	fmt.Println("Conn Established...")

	i := 0
	bytes := make([]byte, 1024)
	for {
		// 传输数据到服务端
		send := fmt.Sprintf("data - %d", i)
		_, err = conn.Write([]byte(send))
		if err != nil {
			fmt.Printf("write failed, err:%v\n", err)
			break
		}

		_, err := conn.Read(bytes)
		if err != nil {
			fmt.Printf("read failed, err:%v\n", err)
			break
		}
		fmt.Printf("client recieve data is %s -%d\n", string(bytes), i)
		time.Sleep(1 * time.Second)
		i++
	}
}

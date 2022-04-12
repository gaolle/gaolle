// package main
//
// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"network"
// 	"os"
// 	"strings"
// )
//
// func main() {
// 	conn, err := net.Dial("tcp", network.GetAddress())
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer conn.Close()
//
// 	newReader := bufio.NewReader(os.Stdin)
// 	for {
// 		input, _ := newReader.ReadString('\n')
// 		inputInfo := strings.Trim(input, "\r\n")
// 		if _, err = conn.Write([]byte(inputInfo)); err != nil {
// 			panic(err)
// 		}
// 		buf := [512]byte{}
// 		n, err := conn.Read(buf[:])
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

package main

import (
	"net"
	"network"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", network.GetAddress())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		for i := 0; i < 10; i++ {
			msg := "Hello, Hello. How are you?"
			data, _ := network.Encode(msg)
			conn.Write(data)
		}
		time.Sleep(time.Second)
	}
}

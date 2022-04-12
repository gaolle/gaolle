// // package main
// //
// // import (
// // 	"bufio"
// // 	"fmt"
// // 	"net"
// // 	"network"
// // )
// //
// // func main() {
// // 	listen, err := net.Listen("tcp", network.GetAddress())
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	defer listen.Close()
// //
// // 	for {
// // 		conn, err := listen.Accept()
// // 		if err != nil {
// // 			panic(err)
// // 		}
// // 		go process(conn)
// // 	}
// // }
// //
// // func process(conn net.Conn) {
// // 	defer conn.Close()
// // 	for {
// // 		reader := bufio.NewReader(conn)
// // 		var buf [124]byte
// // 		n, err := reader.Read(buf[:])
// // 		if err != nil {
// // 			panic(err)
// // 		}
// // 		s := string(buf[:n])
// // 		fmt.Println("recv client data: %s", s)
// // 		conn.Write([]byte(s))
// // 	}
// // }
//
package main

import (
	"bufio"
	"fmt"
	"net"
	"network"
)

func main() {
	listen, err := net.Listen("tcp", network.GetAddress())
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	// defer conn.Close()
	// reader := bufio.NewReader(conn)
	// var buf [1024]byte
	// for {
	// 	_, err := reader.Read(buf[:])
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	s := string(buf[:])
	// 	fmt.Println("recv client data：%s", s)
	// }

	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := network.Decode(reader)
		if err != nil {
			panic(err)
		}
		fmt.Println("recv clinet data: %s", msg)
	}
}

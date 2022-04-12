package network

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"net"
)

var address = "127.0.0.1:6060"

func GetAddress() string {
	return address
}

func GetTCPAddr() *net.TCPAddr {
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", address)
	return tcpAddr
}

type Package struct {
	len  int
	data []byte
}

func Encode(msg string) ([]byte, error) {
	l := int32(len(msg))
	pkg := new(bytes.Buffer)
	if err := binary.Write(pkg, binary.LittleEndian, l); err != nil {
		panic(err)
	}
	if err := binary.Write(pkg, binary.LittleEndian, []byte(msg)); err != nil {
		panic(err)
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lenMsg, err := reader.Peek(4)
	if err != nil {
		panic(err)
	}
	newBuffer := bytes.NewBuffer(lenMsg)
	
	var l int32
	if err = binary.Read(newBuffer, binary.LittleEndian, &l); err != nil {
		panic(err)
	}

	pkg := make([]byte, 4+l)
	if _, err = reader.Read(pkg[:]); err != nil {
		panic(err)
	}
	return string(pkg[4:]), nil
}

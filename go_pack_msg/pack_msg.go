package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
)

type Msg struct {
	msg_len  uint32
	msg_id   uint32
	msg_data []byte
}

func checkerr(err error) {
	if err != nil {
		fmt.Println("数据写入与读取失败")
	}
}

// 封包函数
func PackMsg(len uint32, id uint32, data []byte) ([]byte, error) {
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}

	// 获取一个存放bytes的缓冲区，存储字节序列
	dataBuff := bufferPool.Get().(*bytes.Buffer)
	// 将数据长度写入字节流
	err := binary.Write(dataBuff, binary.LittleEndian, len)
	checkerr(err)

	// 将id写入字节流
	err = binary.Write(dataBuff, binary.LittleEndian, id)
	checkerr(err)

	// 将数据内容写入字节流
	err = binary.Write(dataBuff, binary.LittleEndian, data)
	checkerr(err)

	return dataBuff.Bytes(), nil
}

// 解包函数
func UnpackMsg(data []byte) (*Msg, error) {
	// 创建一个 io.Reader
	boolBuffer := bytes.NewReader(data)
	msg := &Msg{}

	// 读取数据
	err := binary.Read(boolBuffer, binary.LittleEndian, &msg.msg_len)
	checkerr(err)

	err = binary.Read(boolBuffer, binary.LittleEndian, &msg.msg_id)
	checkerr(err)

	// todo: 有问题 读取固定长度
	err = binary.Read(boolBuffer, binary.LittleEndian, &msg.msg_data)
	checkerr(err)

	return msg, nil
}

func main() {
	var msg_data = []byte("hell world")
	bin, err := PackMsg(20, 1001, msg_data)
	checkerr(err)
	fmt.Println(bin)

	msg, err := UnpackMsg(bin)
	checkerr(err)
	fmt.Println(msg.msg_len)
	fmt.Println(msg.msg_id)
	fmt.Println(msg.msg_data)
}

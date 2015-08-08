package mysql

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

type PacketIO struct {
	rb *bufio.Reader
	wb *bufio.Writer

	Sequence uint8
}

// 将net.Conn转换成为 Reader/Writer
// 交互上语义更加明确
func NewPacketIO(conn net.Conn) *PacketIO {
	p := &PacketIO{
		rb: bufio.NewReaderSize(conn, 2048),
		wb: bufio.NewWriterSize(conn, 2048),
	}
	return p
}

//
// 从conn中读取一个Packet， 不包含前面的Header数据
//
func (p *PacketIO) ReadPacket() ([]byte, error) {
	header := []byte{0, 0, 0, 0}

	// 首先读取4个字节
	// 前三个为: Package长度
	if _, err := io.ReadFull(p.rb, header); err != nil {
		return nil, err
	}

	length := int(uint32(header[0]) | uint32(header[1])<<8 | uint32(header[2])<<16)
	if length < 1 {
		return nil, fmt.Errorf("invalid payload length %d", length)
	}

	// sequence这个是什么概念呢?
	// 必须和当前p的Sequence一致
	sequence := uint8(header[3])
	if sequence != p.Sequence {
		return nil, fmt.Errorf("invalid sequence %d != %d", sequence, p.Sequence)
	}

	p.Sequence++

	// 分配内存(gc?)
	data := make([]byte, length)
	if _, err := io.ReadFull(p.rb, data); err != nil {
		return nil, err
	} else {
		// 如何理解代码呢?
		// length最大也就这样了? 2^24 似乎没得选择
		// 选择就是: length == MaxPayloadLen， 全部都是 FFFFFF
		if length < MaxPayloadLen {
			return data, nil
		}

		// 继续读取下一个Packet?
		var buf []byte
		buf, err = p.ReadPacket()
		if err != nil {
			return nil, err
		} else {
			return append(data, buf...), nil
		}
	}
}

//data already have header
func (p *PacketIO) WritePacket(data []byte) error {
	length := len(data) - 4

	for length >= MaxPayloadLen {
		data[0] = 0xff
		data[1] = 0xff
		data[2] = 0xff

		data[3] = p.Sequence

		if n, err := p.wb.Write(data[:4+MaxPayloadLen]); err != nil {
			return ErrBadConn
		} else if n != (4 + MaxPayloadLen) {
			return ErrBadConn
		} else {
			// 准备下一次写数据
			p.Sequence++
			length -= MaxPayloadLen
			data = data[MaxPayloadLen:]
		}
	}

	// 处理最后一个数据包
	data[0] = byte(length)
	data[1] = byte(length >> 8)
	data[2] = byte(length >> 16)
	data[3] = p.Sequence

	if n, err := p.wb.Write(data); err != nil {
		return ErrBadConn
	} else if n != len(data) {
		return ErrBadConn
	} else {
		p.Sequence++
		return nil
	}
}

func (p *PacketIO) Flush() error {
	return p.wb.Flush()
}

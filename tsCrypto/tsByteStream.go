package tsCrypto

import (
	"bytes"
	"encoding/binary"
	_"fmt"
)

type ByteIo struct {
	data []byte
	pos int
	order binary.ByteOrder
}

func NewByteIoBySize(size int, order binary.ByteOrder)(*ByteIo) {
	ret := &ByteIo {
		data: make([]byte, size),
		pos: 0,
		order: order,
	}
	
	return ret
}

func (this *ByteIo)GetData()([]byte) {
	return this.data
}

func (this *ByteIo)GetSize()(int) {
	return len(this.data)
}

func (this *ByteIo)GetPos()(int) {
	return this.pos
}

func (this *ByteIo)SetPos(pos int) {
	this.pos = pos;
}

func (this *ByteIo)byteData(size int)(*bytes.Buffer) {
	return bytes.NewBuffer(this.data[this.pos:this.pos+size])
}

func (this *ByteIo)ReadByte()(byte) {
	var res byte = 0
	res = this.data[this.pos]
	this.pos += 1
	return res
}

func (this *ByteIo)WriteByte(value byte) {
	this.data[this.pos] = value
	this.pos += 1
}

func (this *ByteIo)ReadInt32()(int32) {
	b_buf := this.byteData(4)
	//b_buf.Reset()
	
	//fmt.Println(b_buf.Bytes())
	
	var res int32 = 0
	binary.Read(b_buf, this.order, &res)
	this.pos += 4
	return res
}

func (this *ByteIo)WriteInt32(value int32) {
	b_buf := bytes.NewBuffer([]byte{})
	
	binary.Write(b_buf, this.order, &value)
	//fmt.Println(b_buf.Bytes())
	b_buf.Read(this.data[this.pos:4])
	this.pos += 4
}
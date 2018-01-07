package ioZmq

import (
	"github.com/pebbe/zmq4"
)

const (
	Zmq_REQ    = zmq4.REQ
	Zmq_REP    = zmq4.REP
	Zmq_DEALER = zmq4.DEALER
	Zmq_ROUTER = zmq4.ROUTER
	Zmq_PUB    = zmq4.PUB
	Zmq_SUB    = zmq4.SUB
	Zmq_XPUB   = zmq4.XPUB
	Zmq_XSUB   = zmq4.XSUB
	Zmq_PUSH   = zmq4.PUSH
	Zmq_PULL   = zmq4.PULL
	Zmq_PAIR   = zmq4.PAIR
	Zmq_STREAM = zmq4.STREAM
)

type ZmqCtx struct {
	ctx *zmq4.Context
}

func NewZmqCtx(thread_count int)(*ZmqCtx, error){
	res:=&ZmqCtx{
	}
	var err error
	res.ctx,err=zmq4.NewContext()
	if err!=nil {
		return nil, err
	}
	res.ctx.SetIoThreads(thread_count)
	return res,nil
}

func (this *ZmqCtx)Term() {
	this.ctx.Term()
}

func (me *ZmqCtx)NewZmqSock(t zmq4.Type)(*ZmqSock, error) {
	res:=&ZmqSock{
		
	}
	var err error
	res.sock,err=me.ctx.NewSocket(t)
	if err!=nil {
		return nil, err
	}
	return res, nil
}
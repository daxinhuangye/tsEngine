package ioZmq

import (
	"github.com/pebbe/zmq4"
)

type ZmqSock struct {
	sock *zmq4.Socket
}

func (this *ZmqSock)Info()(string){
	return "zmq"
}

func (this *ZmqSock)Bind(addr string)(error) {
	err:=this.sock.Bind(addr)
	return err
}

func (this *ZmqSock)Connect(addr string)(error) {
	err:=this.sock.Connect(addr)
	return err
}

func (this *ZmqSock)Recv()([]byte, error) {
	msg,err:=this.sock.RecvBytes(zmq4.Flag(0))
	return msg, err
}

func (this *ZmqSock)Send(data []byte)(int, error) {
	size,err:=this.sock.SendBytes(data, zmq4.Flag(0))
	return size, err
}

func (this *ZmqSock)SendById(id []byte, data []byte)(int, error) {
	size,err:=this.sock.SendBytes(id, zmq4.Flag(2))
	if err!=nil {
		return 0, err
	}
	size,err=this.sock.SendBytes(data, zmq4.Flag(0))
	return size, err
}

func (this *ZmqSock)SetIdentity(id string)(error) {
	this.sock.SetIdentity(id)
	return nil
}

func (this *ZmqSock)SetRcvhwm(count int) (error) {
	return this.sock.SetRcvhwm(count)
}

func (this *ZmqSock)SetSndhwm(count int) (error) {
	return this.sock.SetSndhwm(count)
}

func (this *ZmqSock)SetPlainInfo(user_name string, pwd string) (error) {
	this.sock.SetPlainUsername(user_name)
	this.sock.SetPlainPassword(pwd)
	
	return nil
}

func (this *ZmqSock)Close() {
	this.sock.Close()
}
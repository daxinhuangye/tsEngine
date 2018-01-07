package tsCb

import (
	"gopkg.in/couchbaselabs/gocb.v1"
	"time"
	
	"fmt"
)

const (
	No_key int = 1
	Have_key int = 2
)

type CbObj struct {
	 cl *gocb.Cluster
	 bk *gocb.Bucket
	 
	 cbAddr string
	 bucketName string
	 pwd string
}

func NewCbObj(cb_addr string, bucket string, pwd string)(*CbObj, error) {
	res:=&CbObj{
		cbAddr: cb_addr,
		bucketName: bucket,
		pwd: pwd,
	}
	var err error
   	res.cl, err=gocb.Connect(res.cbAddr)
   	if err!=nil {
   		return nil,err
   	}
  	res.bk, err=res.cl.OpenBucket(res.bucketName, res.pwd)
   	if err!=nil {
   		return nil,err
   	}
   	res.cl.SetConnectTimeout(time.Millisecond * 500)
   	//res.bk.SetOperationTimeout(time.Millisecond * 1000)
   	
	return res,nil
}

func (me* CbObj)Insert(key string, value interface{}, t uint32)(error) {
	_,err:=me.bk.Insert(key, &value, t)
	return err
}

func (me* CbObj)Del(key string) (error) {
	_,err:=me.bk.Remove(key, 0)
	return err
}

func (me* CbObj)Update(key string, value interface{}, t uint32, can_insert bool) (error) {
	if can_insert {
		_,err:=me.bk.Upsert(key, &value, t)
		return err
	}

	_,err:=me.bk.Replace(key, &value, 0, t)
	return err
}

func (me* CbObj)Get(key string) (interface{}, error) {
	var value interface{}
	_, err:=me.bk.Get(key, &value)
	if err!=nil {
		return nil, err
	}
	return value, err
}

func (me *CbObj)Lock(key string, t uint32)(error) {
	lock_key := fmt.Sprintf("lock##20151022_key_%s", key)
	err:=me.Insert(lock_key, &key, t)
	return err
}

func (me *CbObj)Unlock(key string) {
	lock_key := fmt.Sprintf("lock##20151022_key_%s", key)
	me.Del(lock_key)
}
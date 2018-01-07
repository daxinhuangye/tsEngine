package tsCb

import (
	"sync"
	"errors"
)

var gTsCbPool *TsCbPool

type TsCbPool struct {
	cbObjs []*CbObj
	buzy []bool
	
	cbAddr string
	bucketName string
	pwd string
	count int
	
	lock sync.Mutex
}

func InitTsCbPool(cb_addr string, bucket string, pwd string, count int)(error) {
	ret := &TsCbPool{
		cbObjs: make([]*CbObj, count),
		buzy: make([]bool, count),
		cbAddr: cb_addr,
		bucketName: bucket,
		pwd: pwd,
		count: count, 
	}
	
	for i:=0; i<count; i++ {
		obj,err := NewCbObj(ret.cbAddr, ret.bucketName, ret.pwd)
		if err!=nil {
			return err
		}
		ret.cbObjs[i] = obj
		ret.buzy[i] = false
		
		obj.Get("test...test...")
	}
	
	gTsCbPool = ret
	
	return nil
}

func getCb()(*CbObj) {
	gTsCbPool.lock.Lock()
	defer gTsCbPool.lock.Unlock()
	for i:=0; i<len(gTsCbPool.buzy); i++ {
		if !gTsCbPool.buzy[i] {
			gTsCbPool.buzy[i] = true
			return gTsCbPool.cbObjs[i]
		}
	}
	return nil
}

func giveBackCb(cb *CbObj) {
	gTsCbPool.lock.Lock()
	defer gTsCbPool.lock.Unlock()
	for i:=0; i<len(gTsCbPool.buzy); i++ {
		if gTsCbPool.cbObjs[i] == cb {
			gTsCbPool.buzy[i] = false
			return 
		}
	}
}

func Insert(key string, value interface{}, t uint32)(error) {
	cb := getCb()
	if cb==nil {
		return errors.New("no free db")
	}
	defer giveBackCb(cb)
	
	err:=cb.Insert(key, value, t)
	return err
}

func Del(key string) (error) {
	cb := getCb()
	if cb==nil {
		return errors.New("no free db")
	}
	defer giveBackCb(cb)
	
	err:=cb.Del(key)
	return err
}

func Update(key string, value interface{}, t uint32, can_insert bool) (error) {
	cb := getCb()
	if cb==nil {
		return errors.New("no free db")
	}
	defer giveBackCb(cb)

	err:=cb.Update(key, value, t, can_insert)
	return err
}

func Get(key string) (interface{}, error) {
	cb := getCb()
	if cb==nil {
		return "", errors.New("no free db")
	}
	defer giveBackCb(cb)
	
	return cb.Get(key)
	//return "0000", nil
}

func GetValues(keys []string)([]interface{}, []error) {
	if len(keys) <= 0 {
		return nil, nil
	}
	cb := getCb()
	if cb==nil {
		return nil, []error{errors.New("no free db")}
	}
	defer giveBackCb(cb)
	
	values := make([]interface{}, len(keys))
	errors := make([]error, len(keys))
	
	for i:=0; i<len(keys); i++ {
		values[i], errors[i] = cb.Get(keys[i])
	}
	
	return values, errors
}